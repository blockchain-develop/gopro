package temp

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

func bytesReverse(u []byte) []byte {
	for i, j := 0, len(u)-1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return u
}

var bigOne = big.NewInt(1)

func BigIntFromNeoBytes(ba []byte) *big.Int {
	res := big.NewInt(0)
	l := len(ba)
	if l == 0 {
		return res
	}

	bytes := make([]byte, 0, l)
	bytes = append(bytes, ba...)
	bytesReverse(bytes)

	if bytes[0]>>7 == 1 {
		for i, b := range bytes {
			bytes[i] = ^b
		}

		temp := big.NewInt(0)
		temp.SetBytes(bytes)
		temp.Add(temp, bigOne)
		bytes = temp.Bytes()
		res.SetBytes(bytes)
		return res.Neg(res)
	}

	res.SetBytes(bytes)
	return res
}

func TestParse(t *testing.T) {
	a := decimal.Decimal{}
	b := decimal.NewFromFloat(0.1)
	c := decimal.NewFromFloat(-0.1)
	fmt.Printf("%d, %d\n", a.Cmp(b), a.Cmp(c))
}
