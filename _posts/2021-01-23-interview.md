---
layout: post 
title: go面试【1】-基础
date: 2021-1-23
categories: interview
tags: [面试]
description: 文章金句。
---
### 1、go的调度 ###
Go基于GMP来调度，在程序运行开始时，n个cpu启动n个线程，之间相互绑定，此时一个线程启动，在经历了函数地址、函数起始地址、参数长度等信息更新后，他会去寻找一个P来绑定，进入M之中运行，如果此时另一个G运行，他会去寻找空闲的P绑定，进入M之中运行， 如果没有空闲的P，则进入全局队列等待M来运行。
### ２、go struct能不能比较 ###
因为是强类型语言，所以不同类型的结构不能比较，但是同一类型的实例值是可以比较的，
### ３、go defer（for defer） ###
先进后出，后进先出,如果存在pannic 则最后执行panic
### ３、select可以用于什么 ###
常用于gorotine的完美退出，用于监听IO操作，当IO操作发生时，触发相应的动作每个case语句里必须是一个IO操作，确切的说，应该是一个面向channel的IO操作
### ５、context包的用途 ###
context通常被译作上下文，它是一个比较抽象的概念，其本质是存在上下层的传递，上层吧内容传递给下层，在go语言中 程序单元也指的是goroutine
### ６、client如何实现长连接###
server是设置超时超时时间，for循环遍历的
### ７、主协程如何等其余协程完再操作 ###
可以使用sync.watigroup、channel、context的WithCancel、select
### 8、slice，len，cap，共享，扩容 ###
切片拥有长度和容量，可以使用make初始化，初始化过后可以使用append函数添加，因为切片底层数据结构是由数组、cap、len组成，所有使用append时，会查看数组后面有没有连续内存块，有就在后面添加，没有就重新生成一个大的数组，并且扩容的时候在容量小于10000的时候，是一倍的增加，当大于10000的时候是1.25倍增加
### ９、map如何顺序读取 ###
map不支持顺序读取，因为他是无序的，，想要有序读取，则需要他的键有序，可以把键有序的放入数组或者切片当中，遍历切片来有序的读取map
### 10、实现set ###
Go中是不提供Set类型的，Set是一个集合，其本质就是一个List，只是List里的元素不能重复。<br>
Go提供了map类型，但是我们知道，map类型的key是不能重复的，因此，我们可以利用这一点，来实现一个set。那value呢？value我们可以用一个常量来代替，比如一个空结构体，实际上空结构体不占任何内存，使用空结构体，能够帮我们节省内存空间，提高性能<br>
代码详情
[set代码实现](https://gitee.com/sixnine/interview/blob/master/set/set.go)
![img](https://deathlimbo.github.io/img/set/1.jpg)
### 11、实现消息队列（多生产者，多消费者） ###
队列是一种数据结构，先进先出的理论，我们可以基于channel实现 多消费者多生产者
[消息队列实现代码实现](https://gitee.com/sixnine/interview/blob/master/MQ/mq.go)
![img](https://deathlimbo.github.io/img/mq/1.png)
![img](https://deathlimbo.github.io/img/mq/2.jpg)
![img](https://deathlimbo.github.io/img/mq/3.jpg)

### 12、大文件排序 ###
### 13、基本排序，哪些是稳定的 ###
快速排序、希尔排序、堆排序、直接选择排序属于不稳定的排序算法<br>
基数排序、冒泡排序、直接插入排序、拆办插入排序、归并排序是稳定的排序算法<br>
1、所谓排序，就是使一串记录，按照其中的某个或某些关键字的大小，递增或者递减的排列起来的操作。排序算法，就是如何使得记录按照要求排列的方法。排序算法在很多领域都得到相当的重视，尤其是在大量数据的处理方面。一个优秀的算法可以节省大量的资源。<br>
2、排序(sort)是计算机程序设计中的一个重要操作，他的功能是将一个数据元素（或记录）的任意序列，重新排列成一个关键字有序的序列<br>
3、稳定性:一个排序算法是稳定的，就是当有两个相等记录的关键字，且在原本的列表中出现，在排序过的列表中也将会是<br>
### 14、http get跟head ###
Head 跟get的本质一样的，区别在于head 不含有呈现数据，而仅仅是http头部信息，使用场景：欲判断某个资源是否存在，我们通常使用GET，但这里使用head则意义更加明确
### 15、http 401,403 ###
400 请求报文存在语法错误<br>
401 发送的请求需要通过HTTP认证的认证信息<br>
403 请求资源的访问被服务器拒绝<br>
404 表示在服务器没有找到请求的资源<br>
### 16、http keep-alive ###
client发出的http请求头需要添加Connection:keep-alive字段<br>
Web-Server端需要识别Connection:keep-alive字段，并且在http的response里指定Connection:keep-alive字段，告诉client，我能提供keep-alive服务，并且应予clinet我暂时不会关闭socket链接
### 17、http能不能一次连接多次请求，不等后端返回 ###
http本质上是使用socket链接，因此发送请求，直接写入tcp缓存，是可以多次进行的，这也是http无状态的原因
### 18、tcp与udp区别，udp优点，适用场景 ###
tcp传输的是数据流，而udp传输的是数据包，tcp会经过三次握手，四次挥手，而udp不会
### 19、time-wait的作用 ###
time-wait开始的时间为tcp四次挥手中主动关闭连接方发送完最后一次挥手，也就是ack=1的信号结束后，主动关闭连接放所处的状态
### 20、数据库如何建索引 ###
1、普通索引<br>
这是最基本的索引，没有任何限制，他有以下创建方式:<br>
(1)create index name on mytable(username(lenght))<br>
如果是char，vachar类型，length可以小于字段实际长度；如果是BLOB和TEXT类型必须指定length<br>
(2)修改表结构<br>
alter mytable addindex[indexname] on (username(lenght))<br>
和上面一致<br>
(3)创建表的时候指定<br>
CREATE TABLE mytable( IDINTNOT NULL, username VARCHAR(16) NOT NULL, <br>
INDEX [indexName] (username(length)) )<br>
(4)删除索引<br>
drop index [indexname]on mytrable<br>
2、唯一索引<br>
索列值必须唯一，但允许有空值。如果是如果是组合索引，则列值的组合必须唯一<br>
(1)创建索引<br>
create unique index indexName on mytable(username(length))<br>
(2)修改表结构<br>
ALTERmytable ADDUNIQUE [indexName] ON (username(length))<br>
(3)创建表的时候直接指定<br>
CREATE TABLE mytable( IDINTNOT NULL, username VARCHAR(16) NOT NULL, <br>
UNIQUE [indexName] (username(length)) ); <br>
3、主键索引<br>
他是一种特俗的唯一索引，不允许有空值。一般是在建表的时候同时创建逐渐索引<br>
create table mytable (<br>
 ID int not null,<br>
 username vachar(16) not null.<br>
 primary key(Id);<br>
)<br>
记住：一个表只有一个主键<br>
4、组合索引<br>
为了形象的比对比单列索引，为表添加多个字段<br>
create table mytable(<br>
    id int not null,<br>
    username vachar(16) not null,<br>
    city vachar(50) not null,<br>
    age int not null,<br>
);<br>
为了进一步榨取mysql的效率，就要考虑建立组合索引，就是将name，city，age见到一个索引里：<br>
alter table mytable addindex name_city_age(name(10),city,age)<br>
建表时，username长度为16，这里用10.这是因为一般情况下名字不的长度不会超过10，这样会加速索引查询速度，还会减少索引文件的速度，增加insert的速度<br>
### 21、孤儿进程，僵尸进程 ###
孤儿进程:一个父进程退出，还有一个或者多个子进程还在运行，那么这些子进程将成为孤儿进程。孤儿进程将被init进程（进程号为1）所收养，并由init进程完成对他们完成状态收集工作。<br>
僵尸进程:一个进程使用fork创建子进程，如果子进程退出，而父进程并没有调用wait或waitpid获取子进程的状态信息，那么子进程的进程描述符任然保存在系统中。<br>
### 22、死锁条件，如何避免 ###
死锁:指多个进程因竞争资源而造成的一种僵局，相互等待，如无外力作用，这些进程都将无法向前推进<br>
死锁的四个必要条件:<br>
 1、互斥条件:一个资源每次只能被一个进程使用，即在一段时间内某个资源仅为一个进程所占有，若有其他进程请求该资源，则请求进程只能等待，<br>
 2、请求与保持条件：进程已经保持了至少一个资源，但又提出了新的资源请求，而该资源已被其他进程占有，此时请求进程被阻塞，但对自己获得的资源保持不放。<br>
 3、不可剥夺条件:进程所获得的资源在未使用完毕之前，不可被其他进程强行夺走，即只能由获得该资源的进程自己来释放<br>
 4、循环等待条件：若干进程间形成首尾相接循环等待资源的关系.<br>
死锁的避免:<br>
 1、破坏不可剥夺条件：一个进程不能获得所需要的全部资源时便处于等待状态，等待期间他占有的资源将被隐式的释放重新加入到资源列表,可以被其他进程使用，而等待的进程只有重新获得自己原有的资源以及新申请的资源才可以重新启动，执行.<br>
 2、破坏“请求与保持条件”:第一种方法静态分配即每个进程在开始时就申请他所需要的全部资源，第二种就是动态分配即每个进程在申请所需要的资源时它本身不占用系统资源<br>
 3、破坏“循环等待”条件：采用资源有序分配其基本思是将系统中的所有资源顺序编号，将紧缺的，稀少的采用较大编号，在申请资源时必须按照编号循环进行，一个进程只有获得较小编号的进程才能申请较大编号的进程<br>
### 23、linux命令，查看端口占用，cpu负载，内存占用，如何发送信号给一个进程 ###
一、查看端口是否被占用：<br>
1.lsof使用， lsof -i:port   查看某个端口是否被占用<br>
2.netstat -anp|grep 9001<br>
netstat -ntlp //查看当前所有tcp端口<br>
netstat -ntulp | grep 80 查看所有80端口的使用情况<br>
netstat -an | grep 3306 查看所有3306的使用情况<br>
参数说明:-t(tcp) 仅显示tcp的相关选项<br>
        -u(udp) 仅显示udp相关选项<br>
        -n 拒绝显示别名,能显示数字的全部转化为数字<br>
        -l 仅列出在listen的服务状态<br>
        -p 显示建立相关链接的程序名<br>
二、cpu负载<br>
使用top查看,<br>
三、内存占用<br>
1、cat /proc/meminfo<br>
2、atop<br>
3、free -h<br>
四、如何发送信号给一个进程<br>
kill 信号参数 进程PID<br> 
### 24、git文件版本，使用顺序，merge跟rebase ###
git文件版本：git -version
### 25、Slice与数组区别，Slice底层结构 ###
slice的底层其实就是数组，只是他的结构体是由数组+容量+长度组成，并且切片是引用传递，数组是值传递
### 26、Go的反射包怎么找到对应的方法 ###

### 27、Redis基本数据结构 ###
redis 有五种基本数据结构:String(字符串)、list(列表)、set(集合)、hash(哈希)、zset(有序集合)
### 28、Redis的List用过吗？底层怎么实现的？ ###
List使用两种数据结构作为底层实现:<br>
 压缩列表ziplist<br>
 双向链表list<br>
 因为双向链表占用的内存比压缩列表要多，所以当创建新的列表键是，列表会优先考虑使用压缩列表，并且在需要的时候转换为双向链表。<br>
 转换条件:<br>
 1、试图往列表添加一个字符串值，且这个字符串的长度超过server.list_max_ziplist_value（默认为64）<br>
 2、ziplist包含节点超过server.list_max_ziplist_entries(默认512)<br>
### 29、Mysql的索引有几种，时间复杂度 ###
mysql有主键索引，唯一索引，普通索引，联合索引,hash索引<br>
b+tress时间复杂度:O(log2 (N-1))<br>
hash索引时间复杂度:O(1)<br>
### 30、InnoDb是表锁还是行锁，为什么 ###
行锁,因为innodb的主索引结构上，既存储了主键值，又直接存储了行数据，可以方便的锁住行数据，而mylsam索引指向另一片数据文件，没有办法精确锁住数据段<br>
### 31、Go的channel ###
channel的底层数据结构式：<br>
 1、环形队列<br>
 2、两个链表实现等待队列<br>
 3、一个互斥锁<br>
 channel基于csp理论的提出，它提供了一种通信机制，可以从一个goroutine向另一个goroutine发送消息，他分为缓冲以及无缓冲状态。可以使用make来创建，配合select来使用<br>

[github上优秀的go面试题总结](https://github.com/lifei6671/interview-go)
[参看文献1](https://www.jianshu.com/p/6bf41d9dcb6e)





