package lock

import "sync/atomic"

type Mutex struct {
	state int32  //是否为唤醒状态
	sema  uint32 //信号量，用于锁释放后唤醒下一个goroutine
}

const (
	mutexLocked           = 1 << iota // 锁定状态位
	mutexWoken                        // 唤醒状态位
	mutexStarving                     // 饥饿状态位
	mutexWaiterShift      = iota      // 第4位开始表示等待数
	starvationThresholdNs = 1e6       //当队首goroutine等待时间> starvationThresholdNs表示饥饿
)

//获取互斥锁
func (m *Mutex) new() {
	// 如果mutex当前状态为空闲，则立即获得锁 （Lock为内联函数）
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
	m.lockSlow()
}

func (m *Mutex) lockSlow() {
	var waitStartTime int64 //当前goroutine等待时间 > starvationThresholdNs则把state设为饥饿
	starving := false       //是否饥饿
	awoke := false          //当前gotoutine是否为唤醒
	iter := 0               //自旋转次数
	old := m.state
	for {
		//如果mutex当前状态为锁定状态且不为唤醒状态，则设置state为唤醒状态，
		//并自旋转等待锁的释放，如果自旋转完成仍然没等到锁释放，则进入队列。
		//如果mutex当前状态为饥饿状态unlock会把mutex的所有权直接交给队首gotoutine，所以饥饿模不需要自旋转
		//如果mutex当前状态为非唤醒的，把它设置为唤醒状态，这样做的目的是unlock不需要再去队列中唤醒其它goroutine
		//runtime_canSpin判断是否可以自旋转(由iter次数控制)
		//atomic.CompareAndSwapInt32自旋转的设置state为唤醒状态
		//runtime_doSpin自旋转的等待unlock
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
		new := old
		// 如果是锁定模式或饥饿模式不会偿试获取锁，老老实实排队
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		if old&(mutexLocked|mutexStarving) != 0 {
			//队列中的等待数量+1，之所以大移3位，因为第4位以后表示等待数
			new += 1 << mutexWaiterShift
		}
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			//如果当前goroutine是被唤醒的，则将唤醒状态设为0，下一步将会用此状态更新state状态
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken
		}
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			//state老状态即不是锁定也不饥饿状态，只可能是唤醒状态，所以成功获取锁
			if old&(mutexLocked|mutexStarving) == 0 {
				break
			}
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 { //首次等待记录开始等待时间
				waitStartTime = runtime_nanotime()
			}
			runtime_SemacquireMutex(&m.sema, queueLifo, 1)
			//如果等待时间>starvationThresholdNs，则starving为true
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 {
				//如果当前goroutine已经被唤醒且mutex是饥饿模式，说明已经获取到锁（unlock会直接把锁交给唤醒的队首goroutine）
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				//如果当前goutine不是饥饿状态或者当前gotoutine是队列中的最后一个，则取消mutex的饥饿模式，并设为锁定模式
				if !starving || old>>mutexWaiterShift == 1 {
					delta -= mutexStarving //delta=-11, stage=12, 相加结果为1(锁定模式)
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}
}

//释放互斥锁
func (m *Mutex) Unlock() {
	//如果是单一的锁定状态，则直接释放成功(内联函数)
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		m.unlockSlow(new)
	}
}

func (m *Mutex) unlockSlow(new int32) {
	if (new+mutexLocked)&mutexLocked == 0 {
		throw("sync: unlock of unlocked mutex")
	}
	if new&mutexStarving == 0 { //普通模式
		old := new
		for {
			//如果队列中没有等待者或锁已经被获取或已经被唤醒或处于饥饿模式时，无需唤醒任何goroutine
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			//等待队列-1，并唤醒队首的goroutine，通过m.sema信号量通知
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, false, 1)
				return
			}
			old = m.state
		}
	} else {
		// 饥饿模式则直接把锁交给队首的等待者
		runtime_Semrelease(&m.sema, true, 1)
	}
}
