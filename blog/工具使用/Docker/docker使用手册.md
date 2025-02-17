# docker使用手册

## 如何查询docker镜像以及其tag

在命令行查询镜像

```
docker search ubuntu
```

例如：
```
root@egaotan-VirtualBox:/etc/docker# docker search ubuntu
NAME                                                      DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
ubuntu                                                    Ubuntu is a Debian-based Linux operating sys…   10919               [OK]
dorowu/ubuntu-desktop-lxde-vnc                            Docker image to provide HTML5 VNC interface …   429                                     [OK]
rastasheep/ubuntu-sshd                                    Dockerized SSH service, built on top of offi…   244                                     [OK]
consol/ubuntu-xfce-vnc                                    Ubuntu container with "headless" VNC session…   218                                     [OK]
ubuntu-upstart                                            Upstart is an event-based replacement for th…   109                 [OK]
neurodebian                                               NeuroDebian provides neuroscience research s…   68                  [OK]
1and1internet/ubuntu-16-nginx-php-phpmyadmin-mysql-5      ubuntu-16-nginx-php-phpmyadmin-mysql-5          50                                      [OK]
ubuntu-debootstrap                                        debootstrap --variant=minbase --components=m…   44                  [OK]
nuagebec/ubuntu                                           Simple always updated Ubuntu docker images w…   24                                      [OK]
i386/ubuntu                                               Ubuntu is a Debian-based Linux operating sys…   20
1and1internet/ubuntu-16-apache-php-5.6                    ubuntu-16-apache-php-5.6                        14                                      [OK]
1and1internet/ubuntu-16-apache-php-7.0                    ubuntu-16-apache-php-7.0                        13                                      [OK]
ppc64le/ubuntu                                            Ubuntu is a Debian-based Linux operating sys…   13
1and1internet/ubuntu-16-nginx-php-phpmyadmin-mariadb-10   ubuntu-16-nginx-php-phpmyadmin-mariadb-10       11                                      [OK]
1and1internet/ubuntu-16-nginx-php-5.6                     ubuntu-16-nginx-php-5.6                         8                                       [OK]
1and1internet/ubuntu-16-nginx-php-5.6-wordpress-4         ubuntu-16-nginx-php-5.6-wordpress-4             7                                       [OK]
1and1internet/ubuntu-16-apache-php-7.1                    ubuntu-16-apache-php-7.1                        6                                       [OK]
darksheer/ubuntu                                          Base Ubuntu Image -- Updated hourly             5                                       [OK]
pivotaldata/ubuntu                                        A quick freshening-up of the base Ubuntu doc…   4
1and1internet/ubuntu-16-nginx-php-7.0                     ubuntu-16-nginx-php-7.0                         4                                       [OK]
pivotaldata/ubuntu16.04-build                             Ubuntu 16.04 image for GPDB compilation         2
1and1internet/ubuntu-16-sshd                              ubuntu-16-sshd                                  1                                       [OK]
smartentry/ubuntu                                         ubuntu with smartentry                          1                                       [OK]
1and1internet/ubuntu-16-php-7.1                           ubuntu-16-php-7.1                               1                                       [OK]
pivotaldata/ubuntu-gpdb-dev                               Ubuntu images for GPDB development              1
```

如果想要查询镜像以及镜像tag，从以下网站查询：

[docker镜像查询](https://hub.docker.com/u/library)

## 查询本机已下载的镜像

```
docker images
```

例如：
```
root@egaotan-VirtualBox:/etc/docker# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
hello-world         latest              bf756fb1ae65        4 months ago        13.3kB
```

## 从镜像仓库下载镜像

下载最新版本的镜像
```
docker pull ubuntu
```

下载指定tag的镜像
```
docker pull ubuntu:13.10
```

## 删除本机镜像

```
docker rmi hello-world
```

## 运行程序输出结果

```
docker run ubuntu /bin/echo "Hello world"
```

## 运行程序并且交互

```
docker run -i -t ubuntu /bin/bash
```

使用exit来退出
```
exit
```

## 启动容器后台运行

```
docker run -d ubuntu /bin/sh -c "while true; do echo hello world;sleep 1; done"
```

例如：
```
root@egaotan-VirtualBox:~# docker run -d ubuntu /bin/sh -c "while true; do echo hello world;sleep 1; done"
3bf8a6ff2444d8fcf9c3b7e604cd97e18f6f8948e5b42ce1f046ded9803a370d
```

以这种模式启动容器，会得到一个容器ID，可以通过以下方式来检查容器运行情况：
```
docker ps
```

例如：
```
root@egaotan-VirtualBox:~# docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
3bf8a6ff2444        ubuntu              "/bin/sh -c 'while t…"   24 seconds ago      Up 22 seconds                           epic_ptolemy
```

如果要查询容器的输出：
```
docker logs [容器ID]
```

例如：
```
root@egaotan-VirtualBox:~# docker logs 3bf8a6ff2444
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
hello world
```

如果要停止容器
```
docker stop [容器ID]
```

例如：
```
root@egaotan-VirtualBox:~# docker stop 3bf8a6ff2444
3bf8a6ff2444
root@egaotan-VirtualBox:~# docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
```

进入容器
```
docker exec -it 69d1 /bin/bash
```

-- ubuntu 18.04
!! -p: 端口映射，格式为 host_port:container_port。
docker run -p 20336:20336 -itd --name ubuntu01 ubuntu:bionic
docker run -p 20336:20336 -p 8545:8545 -itd --name poly ubuntu:bionic
docker run -p 10332:10332 -p 10331:10331 -itd --name neo_mainnet ubuntu:bionic

docker exec -it xxxx /bin/bash

-- ubuntu 22.04
docker run -p 5432:5433 -p 8545:8545 -itd  -v /test:/soft --name polygon-demo ubuntu:jammy
docker run -itd  -v /Users/tangaoyuan/Documents/gopath/src/github.com:/github.com --name solana2 ubuntu:jammy


