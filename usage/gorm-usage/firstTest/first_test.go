package firstTest

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
)


func newDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%d sslmode=disable", "tangaoyuan", "123456", "wallet", 5432)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_",
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}

type Address struct {
	ID           int32 `json:"-"`
	WalletName   string
	Chain        string
	Symbol       string `json:"Symbol,omitempty" gorm:"-"` // TODO - temporary ignore
	Index        uint32
	ClientAddrId int32
	Address      string
	Tag          string
	gorm.Model
}

func TestFirst(t *testing.T) {
	var a Address
	db := newDB()
	err := db.Where(&Address{
		Chain:   "stellar",
		Address: "GA45DF5RWPAFBTSFI3VE3HMOLI5ZRXCFM3FUIBWPT7B3JZXWO2AIISVR",
		Tag:     "",
	}).First(&a).Error
	fmt.Printf("%v, %v", a, err)
}
