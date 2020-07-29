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
export GOPATH=/root/gopath
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

## 参考