package BasicTest

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func newDB() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/palette?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func TestMigrate(t *testing.T) {
	db := newDB()
	db.AutoMigrate(&Company{})
	//db.AutoMigrate(&Profile{})
	db.AutoMigrate(&User{})
}

func TestCreate(t *testing.T) {
	db := newDB()
	companys := make([]*Company, 0)
	for i := 0;i < 2;i ++ {
		company := &Company{
			Code:  fmt.Sprintf("company_%d", i),
			Name:  fmt.Sprintf("company_%d", i),
			Users: make([]*User, 0),
		}
		companys = append(companys, company)
		for j := 0;j < 40;j ++ {
			user := &User{
				Code:      fmt.Sprintf("company_%d_user_%d", i, j),
				Name:      fmt.Sprintf("company_%d_user_%d", i, j),
				Age:       int64(j),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				CompanyID: "",
			}
			company.Users = append(company.Users, user)
		}
	}
	db.Create(companys)
}

func TestSave(t *testing.T) {
	db := newDB()
	companys := make([]*Company, 0)
	for i := 0;i < 4;i ++ {
		company := &Company{
			Code:  fmt.Sprintf("company_%d", i),
			Name:  fmt.Sprintf("company_%d", 1000000 + i),
		}
		companys = append(companys, company)
	}
	db.Save(companys)
}
