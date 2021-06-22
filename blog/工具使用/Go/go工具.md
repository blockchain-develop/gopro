# go开发工具

## go-kit

## fperf

fperf是一个压测框架，[git](https://github.com/fperf/fperf)

## 远程请求协议

+ GRPC
+ HTTP + JSON
+ thrift

## 服务中间件

+ [gin](https://github.com/gin-gonic/gin)
+ [beego](https://beego.me/docs/intro/)
+ [go-kit](https://github.com/go-kit/kit)

## 数据库中间件
+ [gorm](https://github.com/go-gorm/gorm)
+ [xorm]()
+ [gormat](https://github.com/airplayx/gormat)

## 服务拦截器

+ log
+ metric
+ tracing
+ circuitbreaker
+ rate-limiter

## gofmt

```
find ./ -name "*.go" | xargs gofmt -w -l
find . | xargs grep -i "\"UsdtAmount\":\"0\""
find . | xargs grep -i "\"DstChainId\":2,\"UsdtAmount\":\"0"
find . | xargs grep -i "\"SrcChainId\":6,\"Hash\":\"bcf39f0edda668c58371e519af37ca705f2bfcbd\",\"DstChainId\":2"


find . | xargs grep -i "\"Hash\":\"bcf39f0edda668c58371e519af37ca705f2bfcbd\",\"DstChainId\":2"

find . | xargs grep -i "\"Hash\":\"aee4164c1ee46ed0bbc34790f1a3d1fc87796668\",\"DstChainId\":7"
```

## 在windows下跨平台编译linux
GOOS=linux;GOARCH=amd64;CGO_ENABLED=0

## beego根据mysql生成models

bee api xxxx -conn="root:root@tcp(localhost:3306)/palette?charset=utf8"

## go插件开发
[go插件开发](https://www.jianshu.com/p/917b159a4be6)

## 如何在go中添加类似c中的宏

[golang交叉编译和条件编译的实际应用](https://zhuanlan.zhihu.com/p/92235251)




