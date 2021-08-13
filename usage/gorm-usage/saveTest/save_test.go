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
	db.AutoMigrate(&TokenCollect{})
}

func TestSave(t *testing.T) {
	db := newDB()
	collects := make([]*TokenCollect, 0, 10)
	for i := 0;i < 10;i ++ {
		collects = append(collects, &TokenCollect{
			Chain:  "ethereum",
			Coin:   "ETH",
			Status: "Init",
			From:   fmt.Sprintf("%d", i),
			To:     fmt.Sprintf("%d", i),
			Value:  decimal.NewFromInt(1),
			Txid:   "",
		})
	}
	db.Save(&collects)
	for i := 0;i < 10;i ++ {
		collect := collects[i]
		fmt.Printf("id: %d, chain: %s, coin: %s, status: %s, from: %s, to: %s, value: %s, create time: %s\n",
			collect.ID, collect.Chain, collect.Coin, collect.Status,  collect.From, collect.To, collect.Value.String(), collect.CreatedAt.String())
	}
}
