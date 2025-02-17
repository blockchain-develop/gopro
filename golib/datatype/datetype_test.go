package datatype

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

func TestIntSize(t *testing.T) {
	var n int
	n = 12
	fmt.Printf("size of int is %d\n", unsafe.Sizeof(n))
}

func TestDataTypeShift(t *testing.T) {
	tt := time.Now().Unix()
	var tt_int int
	tt_int = int(tt)

	newtt := tt << 16
	var newtt_int int
	newtt_int = int(newtt)

	fmt.Printf("time is %x, new type of time is %x\n", tt, tt_int)
	fmt.Printf("time is %x, new type of time is %x\n", newtt, newtt_int)
}

func TestBytes2HexString(t *testing.T) {
	data := make([]byte, 20)
	dataHex := hex.EncodeToString(data)
	fmt.Printf("hex string: %s\n", dataHex)
}

func TestString2Other(t *testing.T) {
	{
		dataStr := "123"
		data, err := strconv.ParseInt(dataStr, 10, 32)
		if err != nil {
			panic(err)
		}
		fmt.Printf("int64 data: %d\n", data)
	}
	{
		dataStr := "123"
		data, err := strconv.ParseUint(dataStr, 10, 32)
		if err != nil {
			panic(err)
		}
		fmt.Printf("uint64 data: %d\n", data)
	}
	{
		dataStr := "123"
		data, err := strconv.Atoi(dataStr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("int data: %d\n", data)
	}
	{
		dataStr := "false"
		data, err := strconv.ParseBool(dataStr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("bool data: %v\n", data)
	}
	{
		dataStr := "1.2345"
		data, err := strconv.ParseFloat(dataStr, 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("float data: %v\n", data)
	}
}

func TestOther2String(t *testing.T) {
	{
		dataOther := int64(123)
		data := strconv.FormatInt(dataOther, 10)
		fmt.Printf("int64 data: %s\n", data)
	}
	{
		dataOther := uint64(123)
		data := strconv.FormatUint(dataOther, 10)
		fmt.Printf("uint64 data: %s\n", data)
	}
	{
		dataOther := int(123)
		data := strconv.Itoa(dataOther)
		fmt.Printf("int data: %s\n", data)
	}
	{
		dataOther := true
		data := strconv.FormatBool(dataOther)
		fmt.Printf("bool data: %s\n", data)
	}
	{
		dataOther := 1.2345
		data := strconv.FormatFloat(dataOther, 'f', 8, 64)
		fmt.Printf("float data: %s\n", data)
	}
}

func TestInt64ToUint64(t *testing.T) {
	a := int64(-1)
	b := uint64(a)
	fmt.Printf("a: %d, b: %d\n", a, b)
}

func TestDecimalDiv(t *testing.T) {
	d1, _ := decimal.NewFromString("4996771978399024063769")
	denom, _ := decimal.NewFromString("1000000000000000000")
	//d2 := d1.Div(denom)
	d2 := d1.DivRound(denom, 32)
	fmt.Printf("%s\n", d2.String())
}

func TestDecode(t *testing.T) {
	aa, _ := hex.DecodeString("f8c69e91e17587c8d4fe7b00000000000000000000000000af331ba8327fbb35b1c4feff000000000100")
	//aa, _ := hex.DecodeString("f8c69e91e17587c80c2b4a8a040000000000000000000000ffffffffffffffffffffffffffffffff0001")
	a1 := aa[0]
	a2 := binary.LittleEndian.Uint64(aa[8:16])
	a3 := binary.LittleEndian.Uint64(aa[16:24])
	a4 := binary.LittleEndian.Uint64(aa[24:32])
	a5 := binary.LittleEndian.Uint64(aa[32:40])
	a6 := aa[33]
	a7 := aa[34]
	fmt.Printf("length: %d, %d %d %d %d %d %d %d", len(aa), a1, a2, a3, a4, a5, a6, a7)
}

func TestDecode1(t *testing.T) {

	type UpdateEquityReq struct {
		Uid          int64           `json:"userId,omitempty"`
		EquityAmount decimal.Decimal `json:"amount,omitempty"`
		Currency     string          `json:"currency,omitempty"`
		EquityAt     uint64          `json:"transactTimeNs"`
	}

	type UpdateEquityMsg struct {
		EventType string `json:"eventType"`
		Data      struct {
			Wallets []UpdateEquityReq `json:"wallets"`
		} `json:"data"`
	}

	body := "{\"eventType\":\"EQUITY_WALLET_EVENT\",\"data\":{\"wallets\":[{\"userId\":11374943,\"currency\":\"eth\",\"amount\":\"1\",\"transactTimeNs\":1676270187554873962},{\"userId\":11356509,\"currency\":\"eth\",\"amount\":\"1\",\"transactTimeNs\":1676270961045811156},{\"userId\":11356508,\"currency\":\"eth\",\"amount\":\"1\",\"transactTimeNs\":1676270961045811156}]}}"

	var msg UpdateEquityMsg
	err := json.Unmarshal([]byte(body), &msg)
	if err != nil {
		panic(err)
	}
}

func TestSxx(t *testing.T) {
	{
		s := new(big.Int).SetBytes([]byte("tr"))
		fmt.Printf("%d\n", s.Uint64())
	}
	{
		s := new(big.Int).SetBytes([]byte("TransferTransferTransfer"))
		fmt.Printf("%d\n", s.Uint64())
	}
	{
		s := binary.LittleEndian.Uint64([]byte("tr"))
		fmt.Printf("%d\n", s)
	}
	{
		s := binary.LittleEndian.Uint64([]byte("TransferTransferTransfer"))
		fmt.Printf("%d\n", s)
	}
}
