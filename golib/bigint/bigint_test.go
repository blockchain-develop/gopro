package bigint

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	aaa := new(big.Int).SetUint64(math.MaxUint64)
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

func TestFloat2Int(t *testing.T) {
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

func TestFloat2String(t *testing.T) {
	xxx, _ := new(big.Float).SetString("0.00000011911111111111111111111")
	yyy := new(big.Float).Mul(xxx, new(big.Float).SetUint64(100000000))
	yyy.SetMode(big.ToPositiveInf)
	zzz, aaa := yyy.Uint64()
	fmt.Printf("data: %d, %d\n", zzz, aaa)
}

type BigInt big.Int

func (b *BigInt) value() string {
	data := (*big.Int)(b)
	return data.String()
}

func TestBigIntRedefine(t *testing.T) {
	data := new(big.Int)
	data.SetString("1000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
	fmt.Printf("value: %s\n", (*BigInt)(data).value())
}

func TestString2BigFloat(t *testing.T) {
	data := new(big.Float)
	aaa, result := data.SetString("")
	if !result {
		fmt.Printf("invalid")
	} else {
		fmt.Printf("aaa: %s", aaa.String())
	}
}

func TestFloat2String2(t *testing.T) {
	decimal.DivisionPrecision = 2
	aaa := new(big.Float).Mul(new(big.Float).SetFloat64(0.3333), new(big.Float).SetInt64(10000))
	bbb, _ := aaa.Int64()
	percent := ""
	if bbb != 0 {
		ccc := decimal.NewFromInt(bbb)
		ddd := ccc.Div(decimal.NewFromInt(100))
		percent = fmt.Sprintf("%s%s", ddd.String(), "%")
	}
	fmt.Printf("%s\n", percent)
}

func TestFloat2String3(t *testing.T) {
	decimal.DivisionPrecision = 2
	ddd := decimal.NewFromFloat(float64(0.0034) * 10000)
	ddd = ddd.Div(decimal.NewFromInt(100))
	percent := fmt.Sprintf("%s%s", ddd.String(), "%")
	fmt.Printf("%s\n", percent)
}

func TestDecimal(t *testing.T) {
	type Test1 struct {
		total decimal.Decimal
	}
	t1 := &Test1{}
	t1.total = t1.total.Add(decimal.NewFromInt(1))
	fmt.Printf("total: %s", t1.total.String())
}

func TestDecimalAdd(t *testing.T) {
	a := "0.019931995999178948"
	b := "0.019978999999706"
	f := "add"
	da, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	db, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	dc := decimal.Decimal{}
	switch f {
	case "add":
		dc = da.Add(db)
	case "sub":
		dc = da.Sub(db)
	}
	fmt.Printf("%s %s %s is %s\n", da.String(), f, db.String(), dc.String())
}

func TestDecimal2(t *testing.T) {
	a := "0.02"
	b := "0.01"
	c := "0.01"
	d := "0.000021000000294"
	da, err := decimal.NewFromString(a)
	if err != nil {
		panic(err)
	}
	db, err := decimal.NewFromString(b)
	if err != nil {
		panic(err)
	}
	dc, err := decimal.NewFromString(c)
	if err != nil {
		panic(err)
	}
	dd, err := decimal.NewFromString(d)
	if err != nil {
		panic(err)
	}
	dr := da.Add(db).Sub(dc).Sub(dd)
	fmt.Printf("%s\n",dr.String())
}

func TestTemp(t *testing.T) {
	value := new(big.Int).Mul(
		new(big.Int).SetInt64(23456723),
		new(big.Int).Exp(new(big.Int).SetInt64(10), new(big.Int).SetInt64(int64(1)), new(big.Int).SetInt64(10)),
	)
	fmt.Printf("%s", value.String())
}


func TestTemp1(t *testing.T) {
	value := new(big.Int).Div(big.NewInt(1), big.NewInt(2))
	fmt.Printf("value: %s\n",value.String())
}