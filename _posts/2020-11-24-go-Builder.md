---
layout: post
title: go-设计者模式
date: 2020-11-24
categories: go
tags: [设计值模式 golang]
description: 文章金句。
---
# 谈心 #
我们的gin框架最近搭建的差不多了，我们的web-api集成了gin、gorm、websocket以及很多好用的插件，但是对于他们的具体实现以及掌握并不牢固，所以今天写一期gorm学习笔记增强记忆<br>

# gorm #
## 功能展示  ##
 ·全功能ORM（几乎） <br>
 ·
 
 关联（包含一个，包含多个，属于，多对多，多种包含） <br>
 ·Callbacks（创建/保存/更新/删除/查找之前/之后） <br>
 ·预加载（急加载） <br>
 ·事务 <br>
 ·复合主键 <br>
 ·SQL Builder <br>
 ·自动迁移 <br>
 ·日志 <br>
 ·可扩展，编写基于GORM回调的插件 <br>
 ·每个功能都有测试 <br>
 ·开发人员友好 <br>

  <h4>代码实现</h4>
[gin框架目录](https://gitee.com/sixnine/go-webApi.git)
#### 后续 ####
目前暂时实现了 聊天室，独立聊天，请期待下一期websocket
摘自[gin框架搭建](https://studygolang.com/subject/194)，如有侵犯，请联系本人，此篇文章只作为学习记录