package beego_orm

import (
	"github.com/astaxie/beego/orm"
	"testing"
)

func TestBasic(t *testing.T) {
	o := orm.NewOrm()
	o.Using("palette")

	user := new(User)
	user.Hash = "xxxx"
	//o.Read(user)
	o.ReadOrCreate(user, "Hash")
}

func TestForeignKey(t *testing.T) {
	o := orm.NewOrm()
	o.Using("palette")

	profile := new(Profile)
	profile.Aaaa = "1111"
	profile.UserHash = "bbbbb"
}
