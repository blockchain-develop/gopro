# MySql Use Guide

## 安装(ubuntu)

```
apt-get install mysql-server
```

```
mysql_secure_installation
```

## 问题集锦

1. ubuntu上非root用户使用客户端连接数据库报错 "MysqlAccess denied for user root@localhost"

这是由于密码策略或者登录策略配置不对，使用root用户登录数据库，执行以下操作：
```
use mysql;
update user set plugin = "mysql_native_password";
update user set authentication_string = password('root') where user = "root" and Host = "localhost";
flush privileges;
```
重启mysql
```
service mysql restart
```