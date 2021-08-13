package saveTest

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TokenCollect struct {
	Chain  string `gorm:"uniqueIndex:idx_token_collect_chain_address_coin"`
	Coin   string `gorm:"uniqueIndex:idx_token_collect_chain_address_coin"`
	Status string
	From   string `gorm:"column:sender;uniqueIndex:idx_token_collect_chain_address_coin"`
	To     string `gorm:"column:receipt"`
	Value  decimal.Decimal
	Txid   string
	gorm.Model
}
