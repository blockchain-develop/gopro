package datatype

import (
	"encoding/hex"
	"fmt"
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

type testSlice   []int

func TestSliceLen(t *testing.T) {
	xx := make(testSlice, 0)
	xx = append(xx, 0)
	xx = append(xx, 1)
	fmt.Printf("len of slice: %d\n", len(xx))
}
