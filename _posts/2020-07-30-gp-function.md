---
layout: post
title: go-函数总结(持续更新)
date: 2020-07-30
categories: 函数总结
tags: [go,函数总结]
description: 文章金句。
---
函数总结:
<h2>strconv包</h2>
- strconv.Itoa(i int)： 将int类型转换为string,函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字 <br>
- strconv.Atoi(i string):将string类型转换为Int <br>
- string() 和 strconv.Itoa(i int)的区别 <br>
 1、strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字。<br>
 2、string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符，当且仅当data为[]byte类型时string(data)。<br>

<h2>fmt包</h2>
 fmt方法可以大致分为print,scan两类，根据基础方法构建特定方法.<br>
 - print将参数写入字符串或io.writer<br>
 - scan从字符串或io.Reader读取指定数据，并输出<br>
 ###print###
 - 基础模式 [name] print返回默认格式化的字符串<br>
 - 写入模式 F[name] 返回写入字节 ，例如: Fprint<br>
 - 字符模式 S[name] 返回字符 例如:Sprint<br>
 - 模板模式 [bane]f 根据模板格式化 例如:Printf<br>
 - 换行模式 [name]ln 输出后带换行 例如:Println<br>
 
 模式可以互相组合,例如Fprintf,Sprintf,eg:<br>
 默认格式化，返回字符串:fmt.Printf("name: %s", "coco") -----> name23 [115 104 111 119 32 109 101] <br>
 
###Scan###
- Scan 方法的几种模式， 所有方法都返回 写入字节数(n)及错误(err) <br>
- 基础模式: [name] Scan 将输入值写入参数中 <br>
- 读取模式: F[name] 从io.Reader 读取数据 例如: Fscan <br>
- 字符模式: S[name] 从字符串读取数据 例如: Sscan <br>
- 模板模式: [name]f 根据模板提取数据 例如: scanf <br>
- 换行模式: [name]ln 以换行符号做读取结束 例如: scanln <br>






