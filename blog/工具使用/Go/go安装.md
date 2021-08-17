# Go的安装

## 安装

获取Go的安装包。[Go 1.14安装包地址](https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz)

解压安装包
```
tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz
```

将执行文件添加到环境变量
```
export PATH=$PATH:/usr/local/go/bin
```

添加Go环境变量
```
export GOPATH=/data/gopath
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

## Goland的安装

+ 获取Goland 2019.03的安装包。[Goland 2019.03](https://www.jetbrains.com/go/download/other.html)

+ 安装或者解压。

+ 选择试用。

+ 菜单 -> Help -> Edit Custom VM Options, 在文件后面添加"-javaagent:/Users/yaoyao/.jetbrains/jetbrains-agent-v3.0.0.jar=offline"。
"/Users/yaoyao/.jetbrains/jetbrains-agent-v3.0.0.jar"为本机保存的jetbrains-agent-v3.0.0.jar文件。


## goland编译linux版本

GOARCH=amd64;GOOS=linux

## Goland 2021.1最新安装

+ 从官网获取Goland最新版2021系列
+ 安装或者解压
+ 选择使用
+ 从网盘下载ide-eval-resetter-2.2.3.zip，拖放入Goland来安装插件
+ help -> eval reset,  插件有几个功能，选择reset则可以重置试用期

## 参考