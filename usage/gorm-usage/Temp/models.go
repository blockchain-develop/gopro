package BasicTest

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Transaction is the transaction detail in chain
type Transaction struct {
	Chain         string          `json:"chain"`
	Coin          string          `json:"currency"`
	TxType        string          `json:"txType"`
	From          string          `gorm:"column:sender" json:"from"`
	To            string          `gorm:"column:receipt" json:"to"`
	Txid          string          `json:"txHash"`
	Value         decimal.Decimal `json:"value"`
	GasUsed       decimal.Decimal `json:"gasUsed"`
	Tag           string          `json:"tag"`
	Status        bool            `json:"status"` // The transaction is successful or not
	BlockHeight   int64           `json:"blockId"`
	BlockHash     string          `json:"blockHash"`
	Confirmations int64           `json:"confirmations"`
	Time          int64           `json:"time"`
	BlockTime     int64           `json:"blockTime"`
	Hex           string          `json:"hex"`
	N             uint32          `json:"n"`
	gorm.Model    `json:"-"`
}

type SyncState struct {
	ID          int32
	Chain       string
	BlockHeight int64
	gorm.Model
}

type Address struct {
	ID           int32 `json:"-"`
	Chain        string
	Symbol       string `json:"Symbol,omitempty" gorm:"-"` // TODO - temporary ignore
	Index        uint32
	ClientAddrId int32
	Address      string
	Tag          string
	gorm.Model
}

type HotWalletAccount struct {
	ID      int32 `json:"-"`
	Chain   string
	Index   uint32
	Address string
	gorm.Model
}

type WithdrawalOrder struct {
	ID              int32  `json:"-"`
	Chain           string // blockchain name/ID
	Coin            string // blockchain native coin name
	ClientRequestId string
	Status          string
	From            string `gorm:"column:sender"`
	To              string `gorm:"column:receipt"`
	Value           decimal.Decimal
	Tag             string `gorm:"-"`
	Txid            string
	gorm.Model
}

type WalletBalance struct {
	Chain   string          `json:"chain"`
	Coin    string          `json:"coin"`
	Balance decimal.Decimal `json:"balance"`
}

type SpendStatus int32

const (
	SPENDABLE SpendStatus = iota
	UTXOLOCKED
	SPENT
)

type UnspentOut struct {
	Chain       string
	BlockHeight int64
	BlockHash   string
	Address     string
	Txid        string
	Value       decimal.Decimal
	Hex         string
	N           uint32
	SpendStatus SpendStatus
	gorm.Model  `json:"-"`
}

type SpentOut struct {
	UnspentOut
	SpentBlockHeight int64
	SpentBlockHash   string
	SpentTxid        string
}
