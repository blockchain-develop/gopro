# Redis的特性

## Redis 2.2

### 框架模式
+ 单线程，基于Reactor模式的事件驱动机制，有inet域和unix域的网络事件，基于链表的timer事件。inet域和unix域的网络主要处理客户端的command请求，timer事件主要周期性的bgsave、状态检查以及执行(如同步检查)。
+ I/O多路复用支持select、epoll、kqueue模式。

### 数据结构类型
+ 底层数据结构支持双向链表(list)、哈希表(dict)、整数数组(intset)、字符串(sds)、skiplist(跳跃表)、ziplist(压缩列表)、zipmap(压缩字典)。
+ redis支持的数据类型hash(采用zipmap或者dict来实现)、list(采用ziplist或者list来实现)、set(采用dict或者intset来实现)、zset(采用dict和skiplist来实现)、string(采用sds来实现)。

### 持久化
+ redis支持的持久化包括数据快照(通过fork一个子进程来导出数据快照，在此期间可以继续处理用户command请求)、日志记录(将用户命令持久化到磁盘)、重写日志(生成可以重新生成整个数据的日志记录)。

### 副本机制
+ 支持master-slave的主备模式，且为弱一致性。master执行完command立即返回结果，同时写日志记录，同步日志记录到slave做副本。
+ 仅仅支持完全重同步模式，master发送当前的数据快照以及后续的日志记录到slave执行同步。
+ slave的机器只执行slave of和info两个command，不执行其他command，也就是说slave机器既不能用于读也不能用于写。

### 应用特性
+ 支持key的过期设置。
+ 支持算法命令sort。