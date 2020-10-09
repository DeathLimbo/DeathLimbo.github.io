---
layout: post
title: gin框架搭建websocket,实现聊天室
date: 2020-10-09
categories: websocket
tags: [websocket,协程 golang]
description: 文章金句。
---
# 致歉 #
距离上次更新博客已经快差了一个月，实在是因为公司上线迭代太快，没有个人时间完成自己的学习，导致博客更新放缓<br>

## 如何实现即时聊天 ##
在如何实现及时聊天之前，我们需要去了解用何种方法实现即时聊天（因为本人也是第一次学习并且有强烈的想法想做一个即时聊天的服务，所以基本都是在网上搜索学习到的），以下为实现即时聊天的几个方法：<br>
### 1、拉模式(定时轮询访问接口获取数据) ###
 优缺点:数据更新频率低，则大多数的数据请求时无效的<br>
      在线用户数量多，则服务端的查询负载很高<br>
      定时轮询拉去，无法满足时效性要求<br>
     
### 2、推模式(向客户端进行数据的推送) ###
  优缺点: 仅在数据更新时，才有推送给<br>
       需要维护大量的在线上链接<br>
       数据更新后，可以立即推送<br>
### 3、基于websocket协议做推送 ###
   优点:浏览器支持的socket编程，轻松维持服务端的长连接<br>
       基于TCP协议之上的高层协议，无需开发者关心通讯细节<br>
       提供了高度抽象的编程接口，业务开发成本较低<br>
       <h4>3.1webSocke协议</h4>
       ![img](https://img-blog.csdn.net/20180811160716209?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L1dpbmdfOTM=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
        1、客户端首先发起一个Http请求到服务端，请求的特殊之处，在于请求里面带了一个upgrade的字段，告诉服务端,我想生成一个websocket的协议。<br>
        2、服务端收到请求后，会给客户端一个握手的确认，返回一个switching，意思允许客户端向websocket协议转换,完成这个协议之后，客户端与服务端之间的底层tcp协议是没有中断的<br>
        3、客户端可以向服务端发起一个基于websocket协议的消息<br>
        4、服务端也可以主动向客户端发起websocket协议的消息<br>
        5、，websocket协议里面通讯的单位就叫message<br>
        <h4>3.2协议传输原理</h4>
        · 协议升级后，继续复用Http协议的底层socket完成后续通讯<br>
        · message底层会被切分成多个frame帧进行传输，从协议层面不能传输一个大包，只能切成一个小包传输<br>
        · 编程时，只需要操作message，无需关心frame(属于协议和类库自身去操作的)<br>
        · 框架底层完成TCP网络/IO，websocket协议的解析，开发者无需关心<br>
### 4、基于GO实现WebSocket服务端 ###
  <h4>轮子的选用</h4>
   · 官方的net包 :go语言自带的网络连接相关包, 其中有IP、TCP连接socket、UDP连接socket 等 <br>
   · golang.org/x/net/websocket ：这个包是go语言内写的websocket， socket是一个短连接，客户端发送完信息就结束连接了，不符合我们的实际应用，于是就有了对socket的处理升级为长连接，也就是全双工通信的理念。 这个websocket包就算go语言自带的长连接，不过我没有具体研究使用，看到很多网友说这个包功能性不完善 <br>
   · github.com/gorilla/websocket : 这个包使用的比较多，在这里我就不多说了，有需要可以去github上去看看：[地址](https://github.com/gorilla/websocket/)  <br>
                                    server 启动以后会注册两个 Handler。 <br>
                                    websocketHandler 用于提供浏览器端发送 Upgrade 请求并升级为 WebSocket 连接。 <br>
                                    pushHandler 用于提供外部推送端发送推送数据的请求。 <br>
                                    浏览器首先连接 websocketHandler （默认地址为 ws://ip:port/ws）升级请求为 WebSocket 连接，当连接建立之后需要发送注册信息进行注册。 <br>
                                    推送端发送数据的请求到 pushHandler（默认地址为 ws://ip:port/push）， <br>
                                    Upgrader *websocket.Upgrader，这是 gorilla/websocket 包的对象，它用来升级 HTTP 请求。 <br>
  <h4>代码实现</h4>
[gin框架目录](https://gitee.com/sixnine/go-webApi.git)
#### 后续 ####
目前暂时实现了 聊天室，独立聊天，请期待下一期websocket
摘自[gin框架搭建](https://studygolang.com/subject/194)，如有侵犯，请联系本人，此篇文章只作为学习记录