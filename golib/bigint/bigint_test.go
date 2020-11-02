package bigint

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	aaa := big.NewInt(math.MaxInt64)
	fmt.Printf("max int 64: %d, string: %s\n", aaa.Int64(), aaa.String())
	bbb := aaa.Mul(big.NewInt(math.MaxInt64), aaa)
	fmt.Printf("max big int: %d, string: %s\n", bbb.Int64(), bbb.String())
}

func TestBigInt2(t *testing.T) {
	total := new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(1000000000000000000))
	price1 := new(big.Int).Set(total)
	price2 := new(big.Int).Set(total)
	//
	aaa := big.NewInt(math.MaxInt64)
	bbb := new(big.Int).Mul(aaa, price1)
	fmt.Printf("aaa string: %s, bbb string: %s\n", aaa.String(), bbb.String())

	//
	ccc := big.NewInt(math.MaxInt64)
	ddd := new(big.Int).Mul(aaa, price2)
	fmt.Printf("ccc string: %s, ddd string: %s\n", ccc.String(), ddd.String())

	//
	eee := new(big.Int).Add(aaa, big.NewInt(1))
	mix := new(big.Int).Mul(ccc, eee)
	mix.Add(mix, aaa)

	//
	fff := new(big.Int).Mod(mix, eee)
	ggg := new(big.Int).Div(mix, eee)
	hhh := new(big.Int).Mul(fff, price1)
	iii := new(big.Int).Mul(ggg, price2)
	fmt.Printf("fff string: %s, hhh string: %s\n", fff.String(), hhh.String())
	fmt.Printf("ggg string: %s, iii string: %s\n", ggg.String(), iii.String())
}

func TestBig2String(t *testing.T) {
	a := big.NewInt(999)
	fmt.Printf("a = %d\n", a)

	b := a.String()
	fmt.Printf("b =  %s\n", b)

	c, _ := new(big.Int).SetString(b, 10)
	fmt.Printf("c = %d\n", c)
}

func TestBigDiv(t *testing.T) {
	a := big.NewInt(999)
	b, _ := new(big.Float).SetString(a.String())
	b.Quo(b, big.NewFloat(float64(10)))
	fmt.Printf("b = %s, %s\n", b.String(), b.Text('f', 2))
	x1 := b.Text('f', 18)
	index := len(x1) - 2
	for ;index >=0;index -- {
		if x1[index] == '0' {
			continue
		} else if x1[index] == '.' {
			index --
			break
		} else {
			break
		}
	}
	x1 = x1[0: index + 1]
	fmt.Printf("xxxxx: %s\n", x1)
}

func TestBigDiv1(t *testing.T) {
	amount, _ := new(big.Int).SetString("111111111100", 10)
	a := decimal.NewFromBigInt(amount, 0)
	precision, _ := new(big.Int).SetString("10000000000000000000", 10)
	b := decimal.NewFromBigInt(precision, 0)
	result := a.Div(b)
	fmt.Printf("xxxxx: %s\n", result.String())
}

func TestFloat2String(t *testing.T) {
	a := big.NewFloat(99)
	b, _ := new(big.Int).SetString(a.String(), 10)
	fmt.Printf("a = %s, b = %d\n", a.String(), b.Int64())
}

func TestXxxx(t *testing.T) {
	amount, err := new(big.Int).SetString("<nil>", 10)
	if err != true {
		fmt.Printf("err:\n")
	} else {
		fmt.Printf("amount: %s\n", amount.String())
	}
}