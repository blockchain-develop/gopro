# Postgresql Use Guide

## install on mac

brew install postgresql

## 后台启动

brew services start postgresql

## 非后台启动

pg_ctl -D /usr/local/var/postgres start

## 常用操作

psql -h 127.0.0.1 -U tangaoyuan wallet   -- 连接数据库
\c wallet    -- 选择数据库
select * from pg_tables where schemaname = 'public'   --  查询数据中所有表
\d t_transactions   -- 查看数据表结构
\dt  --列出所有的表
\l --列出所有数据库
\drop database --删除数据库

## postgresql的migration
[How to write & run database migration in Golang](https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g)

* migrate -path db/migrations -database "postgresql://tangaoyuan:123456@127.0.0.1:5432/wallet?sslmode=disable" -verbose up

## aws的主从

as routing protocol


pg_ctl -D /usr/local/var/postgres start

Stop manually
pg_ctl -D /usr/local/var/postgres stop

Start automatically
"To have launchd start postgresql now and restart at login:"

brew services start postgresql

What is the result of pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start?

What is the result of pg_ctl -D /usr/local/var/postgres status 


