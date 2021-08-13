package saveTest

import (
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func newDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%d sslmode=disable", "tangaoyuan", "123456", "template1", 5432)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db
}

func TestMigrate(t *testing.T) {
	db := newDB()
	db.AutoMigrate(&AddressHold{})
}

func TestSave(t *testing.T) {
	db := newDB()
	collects := make([]*AddressHold, 0, 10)
	for i := 0;i < 10;i ++ {
		collects = append(collects, &AddressHold{
			Chain:  "ethereum",
			Coin:   "ETH",
			Index: uint32(i),
			Address:   Address(fmt.Sprintf("%d", i)),
			Balance:  decimal.NewFromInt(1),
			Freeze: decimal.NewFromInt(1),
		})
	}
	db.Save(&collects)
	//
	ins := make([][]interface{}, 0)
	for i := 0;i < 5;i ++ {
		ins = append(ins, []interface{}{"ethereum", "ETH", fmt.Sprintf("%d", i)})
	}
	selects := make([]*AddressHold, 0)
	db.Where("(chain, coin, address) IN (?)", ins).Find(&selects)
	for _, k := range selects {
		fmt.Printf("%d, %s, %s, %s, %d\n", k.ID, k.Chain, k.Coin, k.Address, k.Index)
	}
}
