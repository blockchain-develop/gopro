package saveTest

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Address string

type AddressHold struct {
	Chain   string `gorm:"uniqueIndex:idx_address_hold_chain_address_coin"`
	Coin    string `gorm:"uniqueIndex:idx_address_hold_chain_address_coin"`
	Index   uint32
	Address Address `gorm:"uniqueIndex:idx_address_hold_chain_address_coin"`
	Balance decimal.Decimal
	Freeze  decimal.Decimal
	gorm.Model
}
