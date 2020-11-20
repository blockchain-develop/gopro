package beego_orm

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Hash   string     `orm:"column(hash);pk"`
	Name   string     `orm:"column(name);size(64)"`
	Age    uint64     `orm:"column(age)"`
}

type Profile struct {
	Aaaa     string    `orm:"column(aaaa);pk"`
	Bbbb     string    `orm:"column(bbbb);size(64)"`
	UserHash *User     `orm:"column(user_hash);rel(fk)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/default?charset=utf8")
	orm.RegisterDataBase("palette", "mysql", "root:root@tcp(localhost:3306)/palette?charset=utf8")
	orm.RegisterModel(new(User), new(Profile))
}
