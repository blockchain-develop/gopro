package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*
CREATE SCHEMA IF NOT EXISTS `mysql_test` DEFAULT CHARACTER SET utf8;
USE `mysql_test`;

DROP TABLE IF EXISTS `multiStatements_test`;
CREATE TABLE `multiStatements_test` (
 `name` VARCHAR(100) NOT NULL,
 `id`  INT(4) NOT NULL,
 `url` VARCHAR(256) NOT NULL,
 `height` INT(4) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `InsertBatch_test`;
CREATE TABLE `InsertBatch_test` (
 `name` VARCHAR(100) NOT NULL,
 `id`  INT(4) NOT NULL,
 `url` VARCHAR(256) NOT NULL,
 `height` INT(4) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `UpdateBatch_test`;
CREATE TABLE `UpdateBatch_test` (
 `name` VARCHAR(100) NOT NULL,
 `id`  INT(4) NOT NULL,
 `url` VARCHAR(256) NOT NULL,
 `height` INT(4) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;











DROP TABLE IF EXISTS `MysqlInsertBatch_test`;
CREATE TABLE `MysqlInsertBatch_test` (
 `name` VARCHAR(100) NOT NULL,
 `id`  INT(4) NOT NULL,
 `url` VARCHAR(256) NOT NULL,
 `height` INT(4) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `MysqlUpdateBatch_test`;
CREATE TABLE `MysqlUpdateBatch_test` (
 `name` VARCHAR(100) NOT NULL,
 `id`  INT(4) NOT NULL,
 `url` VARCHAR(256) NOT NULL,
 `height` INT(4) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `MysqlUpdateBatch_test2`;
CREATE TABLE `MysqlUpdateBatch_test2` (
 `name` VARCHAR(100) NOT NULL,
 `id`  INT(4) NOT NULL,
 `url` VARCHAR(256) NOT NULL,
 `height` INT(4) NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;
*/

var DBUrl string = "localhost:3306"
var DBUser string = "root"
var DBPassword string = "root"
var DBName string = "mysql_test"
var DefDB *sql.DB

func init() {
	db, dberr := sql.Open("mysql", "root:root@tcp("+"localhost:3306)/mysql_test?charset=utf8&multiStatements=true")
	if dberr != nil {
		panic(dberr)
	}
	DefDB = db
}
