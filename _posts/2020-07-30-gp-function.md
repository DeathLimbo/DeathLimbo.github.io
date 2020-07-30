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


摘自[7天学会geecache](https://geektutu.com/post/geecache.html)，如有侵犯，请联系本人，此篇文章只作为学习记录






