# Docker安装手册

## 安装Docker

### docker版本

Docker的旧版本称为docker，docker.io或者docker-engine，如果已安装，先卸载。

```
apt-get remove docker docker-engine docker.io containerd runc
```

当前的新版本的Docker Enging-Community软件包为docker-ce。
支持以下ubuntu版本
+ Xenial 16.04 (LTS)
+ Bionic 18.04 (LTS)
+ Cosmic 18.10
+ Disco 19.04

### 安装依赖项

```
apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common
```

### 设置docker安装源

添加Docker的官方GPG密钥

```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
```

添加docker软件源

```
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu  $(lsb_release -cs) stable"
```

### 安装docker-ce

```
apt-get update
apt-get install docker-ce docker-ce-cli containerd.io
```

### 安装结果测试

```
docker run hello-world
```

结果：
```
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
1b930d010525: Pull complete                                                                                                                                  Digest: sha256:c3b4ada4687bbaa170745b3e4dd8ac3f194ca95b2d0518b417fb47e5879d9b5f
Status: Downloaded newer image for hello-world:latest


Hello from Docker!
This message shows that your installation appears to be working correctly.


To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.


To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash


Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/


For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```

### 安装问题

如果出现以下错误
```
docker: Error response from daemon: Get xxx
docker: error pulling image configuration: Get  xxx
```

是因为Get后面指定的url无法解析，可以通过DNS查询命令Dig来查询域名的ip，在本机/etc/hosts中配置域名解析。


如出现以下错误：
```
root@egaotan-VirtualBox:~# docker run hello-world
Unable to find image 'hello-world:latest' locally
docker: Error response from daemon: Get https://registry-1.docker.io/v2/: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers).
See 'docker run --help'.
```

域名无法解析，使用Dig来查询域名ip:
```
root@egaotan-VirtualBox:~# dig @114.114.114.114 registry-1.docker.io

; <<>> DiG 9.10.3-P4-Ubuntu <<>> @114.114.114.114 registry-1.docker.io
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 2373
;; flags: qr rd ra; QUERY: 1, ANSWER: 8, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 512
;; QUESTION SECTION:
;registry-1.docker.io.		IN	A

;; ANSWER SECTION:
registry-1.docker.io.	35	IN	A	23.22.155.84
registry-1.docker.io.	35	IN	A	34.195.246.183
registry-1.docker.io.	35	IN	A	3.218.162.19
registry-1.docker.io.	35	IN	A	35.174.73.84
registry-1.docker.io.	35	IN	A	18.232.227.119
registry-1.docker.io.	35	IN	A	3.94.35.164
registry-1.docker.io.	35	IN	A	3.224.175.179
registry-1.docker.io.	35	IN	A	3.211.199.249

;; Query time: 17 msec
;; SERVER: 114.114.114.114#53(114.114.114.114)
;; WHEN: Sat May 23 19:50:06 CST 2020
;; MSG SIZE  rcvd: 177
```
查询到域名ip后，编辑本机/etc/hosts来添加本机的域名解析：
```
127.0.0.1       localhost
127.0.1.1       egaotan-VirtualBox
23.22.155.84    registry-1.docker.io
34.195.246.183    registry-1.docker.io
18.213.137.78   auth.docker.io
107.23.149.57   auth.docker.io
104.18.123.25    production.cloudflare.docker.com
104.18.121.25    production.cloudflare.docker.com
34.193.164.221   index.docker.io
50.16.172.3     index.docker.io

# The following lines are desirable for IPv6 capable hosts
::1     ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
```

## 安装Docker Compose

Compose是用于定义和运行多容器Docker应用程序的工具。

下载最新的稳定版本：
```
curl -L "https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

添加文件的执行权限
```
chmod +x /usr/local/bin/docker-compose
```

添加软连接，便于运行
```
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
```

测试安装结果
```
docker-compose --version
```

## 参考
[docker教程](https://www.runoob.com/docker/ubuntu-docker-install.html)

[docker官方安装手册](https://docs.docker.com/get-docker/)

[docker compose官方手册](https://docs.docker.com/compose/)

