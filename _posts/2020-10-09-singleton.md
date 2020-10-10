---
layout: post
title: go-单列模式(singleton)
date: 2020-10-10
categories: go
tags: [singleton,golang]
description: 文章金句。 
---
# 单列模式 #
单列模式:保证一个类只有一个实列，并提供一个访问它的全局访问点<br>

通常让一个全局变量成为对象被访问，但不能防止实例化多个对象，一个最好的方法是，让类自身负责保存他的唯一实例,这个类可以保证没有其他实例可以被创建，并且它可以提供一个访问该实例的方法。<br>


## 适用场景 ##
当某个对象只能存在一个实例的情况下使用单例模式(如mysql的链接)
## 优点 ##
1、提供了对唯一实例的受控访问<br>
2、由于在系统内存中只存在一个对象，因此可以节约系统资源，对于一些需要频繁创建和销毁的对象，单例模式可以提高系统性能<br>
## 缺点 ##
1、由于单例模式没有抽象层，因此单例类的扩展有很大的困难 <br>
2、单例类的职责过重，在一定程度上违背了单一职责原则 <br>
### 饿汉模式 ###
饿汉模式:在程序加载阶段init()创建对象实例，虽然保证线程安全，但是会减慢程序启动速度。如果对象实例不被调用，则会浪费一段内存空间。<br>
type SingletonA struct{} <br>
var instanceA *SingletonA <br>
func init() { <br>
instanceA = &SingletonA{} <br>
}<br>
func GetInstanceA() *SingletonA{ <br>
    return instanceA <br>
} <br>
### 懒汉模式 ###
懒汉模式:在获取对象实例时，如果实例为空则创建，虽然避免饿汉模式的空间浪费，但是在多线程场景下，会创建多个对象实例，所以线程非安全。<br>
type SingletonLazy struct {} <br>
var instanceB *SingletonLazy <br>
func GetInstanceLaze() *SingletonLazy { <br>
	if instanceB == nil { <br>
		instanceB = &SingletonLazy{} <br>
	} <br>
	return instanceB <br>
} <br>
### 双重检查机制（锁） ###
type SingletonC struct {} <br>
var instanceC *SingletonC <br>
var mx sync.Mutex <br>
func GetInstanceC() *SingletonC{ <br>
	mx.Lock() <br>
	defer mx.Unlock() <br>
	if instanceC == nil { <br>
		instanceC = &SingletonC{} <br>
	} <br>
	return instanceC <br>
} <br>
### sync.Once ###
为了解决锁带来的开销，因此使用sync.Once，其本质是通过原子性操作，只有满足一定条件时才进行对象创建。<br>
type SingletonD struct {} <br>
var instanceD *SingletonD <br>
var once sync.Once <br>
func GetInstanceD() *SingletonD{ <br>
	once.Do(func() { <br>
		instanceD = &SingletonD{} <br>
	}) <br>
	return instanceD <br>
} <br>
#### 后续 ####
目前暂时实现了 聊天室，独立聊天，请期待下一期websocket
摘自[gin框架搭建](https://studygolang.com/subject/194)，如有侵犯，请联系本人，此篇文章只作为学习记录