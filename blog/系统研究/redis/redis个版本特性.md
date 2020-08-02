# Redis的特性

## 2.2

+ slave的机器只执行slave of和info两个command，不执行其他command，也就是说slave机器既不能用于读也不能用于写。
+ 单线程，基于Reactor模式的事件驱动机制，有inet域和unix域的网络事件，基于链表的timer事件。inet域和unix域的网络主要处理客户端的command请求，timer事件主要周期性的bgsave、状态检查以及执行(如同步检查)。
+ I/O多路复用支持select、epoll、kqueue模式