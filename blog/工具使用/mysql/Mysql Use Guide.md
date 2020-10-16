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

2. 配置远程登录

mysql默认监听的地址是127.0.0.1，首先需要配置监听地址。

修改/etc/mysql/mysql.conf.d/mysqld.cnf中的配置项，去掉指定的监听地址  "bind-address          = 127.0.0.1".

这时远程访问，还是会出现以下错误
```
ERROR 1130 (HY000): Host '172.168.3.79' is not allowed to connect to this MySQL server
```

这是由于mysql访问控制引起的，MySQL不允许远程登录。

```
 update user set host = '%' where user = 'root';
 flush privileges;
```

重启mysql
```
service mysql restart
```

3.



## mysql数据导入和导出

mysqldump -u root -p --databases cross_chain_explorer > cross_chain_explorer_sql

mysql -u root -p cross_chain_explorer < cross_chain_explorer_sql


alter table table_sample change col_sample col_sample varchar(6);

## mysql导出查询结果到文件

mysql -h 127.0.0.1 -u root -p -e `select * from table` > ./output.csv
