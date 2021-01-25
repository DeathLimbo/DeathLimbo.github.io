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
### ３、go defer（for defer） ###
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
![img](https://deathlimbo.github.io/img/set/1.jpg)
[set实现](https://gitee.com/sixnine/interview/blob/master/set/set.go)
### 11、实现消息队列（多生产者，多消费者） ###
队列是一种数据结构，先进先出的理论，我们可以基于channel实现 多消费者多生产者
![img](https://deathlimbo.github.io/img/mq/1.png)
![img](https://deathlimbo.github.io/img/mq/2.jpg)
![img](https://deathlimbo.github.io/img/mq/3.jpg)
[消息队列实现](https://gitee.com/sixnine/interview/blob/master/MQ/mq.go)
### 12、大文件排序 ###
### 13、基本排序，哪些是稳定的 ###

