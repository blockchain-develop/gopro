package comsort

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIsOrderly1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(1))
}

func TestIsOrderly2(t *testing.T) {
	data := [...]int32{5,4,3,2,1}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(2))
}

func TestIsOrderly3(t *testing.T) {
	data := [...]int32{}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(1))
}

func TestIsOrderly4(t *testing.T) {
	data := [...]int32{5}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(1))
}

func TestIsOrderly5(t *testing.T) {
	data := [...]int32{5, 4}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(2))
}

func TestIsOrderly6(t *testing.T) {
	data := [...]int32{1,1, 1, 1, 1}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(1))
}

func TestIsOrderly7(t *testing.T) {
	data := [...]int32{1,1, 1, 1, 1, 4}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(1))
}

func TestIsOrderly8(t *testing.T) {
	data := [...]int32{1,1, 1, 1, 1, 0}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(2))
}

func TestIsOrderly9(t *testing.T) {
	data := [...]int32{1, 2, 3, 3, 5}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(1))
}

func TestIsOrderly10(t *testing.T) {
	data := [...]int32{5,4,3,3,1}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(2))
}

func TestIsOrderly11(t *testing.T) {
	data := [...]int32{1, 2, 3, 3, 5, 3}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(0))
}

func TestIsOrderly12(t *testing.T) {
	data := [...]int32{5,4,3,3,1,3}
	order := IsOrderly(data[:])
	for i:= 0;i < len(data);i ++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Printf("\n")
	fmt.Printf("order is %d\n", order)

	assert.Equal(t, order, byte(0))
}

