# drop & truncate & delete区别

## 用法和区别

* drop：drop table 表名

  删除内容和定义，并释放空间。执行drop语句，将使此表的结构一起删除。

* truncate (清空表中的数据)：truncate table 表名

  删除内容、释放空间但不删除定义(也就是保留表的数据结构)。与drop不同的是,只是清空表数据而已。

  truncate不能删除行数据，虽然只删除数据，但是比delete彻底，它只删除表数据。

* delete：delete from 表名 （where 列名 = 值）

  与truncate类似，delete也只删除内容、释放空间但不删除定义；但是delete即可以对行数据进行删除，也可以对整表数据进行删除。

## 注意

1. delete语句执行删除的过程是每次从表中删除一行，并且同时将该行的删除操作作为事务记录在日志中保存，以便进行进行回滚操作。
   
2. 执行速度一般来说：drop>truncate>delete
   
3. delete语句是数据库操作语言(dml)，这个操作会放到 rollback segement 中，事务提交之后才生效；如果有相应的 trigger，执行的时候将被触发。
   
4. truncate、drop 是数据库定义语言(ddl)，操作立即生效，原数据不放到 rollback segment 中，不能回滚，操作不触发trigger。
   
5. truncate语句执行以后，id标识列还是按顺序排列，保持连续；而delete语句执行后，ID标识列不连续