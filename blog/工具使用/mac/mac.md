# mac

## 进程与端口

lsof -i tcp:8080 

记录一些mac下的常用命令：

1、查看进程号

ps -ef | grep 进程名

2、查看端口被哪个进程监听

sudo lsof -i :端口

3、查看进程监听的端口

sudo lsof -nP -p 进程号 | grep LISTEN

sudo lsof -nP | grep LISTEN | grep 进程号

4、查看监听端口的进程

sudo lsof -nP | grep LISTEN | grep 端口号

5、看到一个新的方法（MacOS统计TCP/UDP端口号与对应服务）

echo "### TCP LISTEN ###"
lsof -nP -iTCP -sTCP:LISTEN
