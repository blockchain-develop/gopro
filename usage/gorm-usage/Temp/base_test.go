package BasicTest

import (
	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
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
	db.AutoMigrate(&Transaction{}, &SyncState{}, &Address{})
}

func TestSave(t *testing.T) {
	db := newDB()

	db.Save(&Transaction{
		Chain:         "cosmos",
		Coin:          "atom",
		TxType:        "Deposit",
		From:          "cosmos1",
		To:            "cosmos2",
		Txid:          "123456",
		Value:         decimal.NewFromInt(1000),
		GasUsed:       decimal.NewFromInt(10),
		Tag:           "",
		Status:        true,
		BlockHeight:   10,
		BlockHash:     "123456",
		Confirmations: 1,
		Time:          0,
		BlockTime:     0,
		Hex:           "",
		N:             0,
		Model:         gorm.Model{},
	})
}
