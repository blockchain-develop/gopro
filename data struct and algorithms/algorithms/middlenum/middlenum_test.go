package middlenum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestMiddleNum1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	middle := MiddleNum(data[:])
	assert.Equal(t, middle, float32(3))
}

func TestMiddleNum2(t *testing.T) {
	data := [...]int32{1, 2, 3, 4}
	middle := MiddleNum(data[:])
	assert.Equal(t, middle, float32(2.5))
}

func TestMiddleNum3(t *testing.T) {
	data := [...]int32{1, 2, 2, 4}
	middle := MiddleNum(data[:])
	assert.Equal(t, middle, float32(2))
}

func TestMiddleNum4(t *testing.T) {
	data := [...]int32{1, 2, 2, 2, 4}
	middle := MiddleNum(data[:])
	assert.Equal(t, middle, float32(2))
}

func TestMiddleNum5(t *testing.T) {
	data := [...]int32{1, 2}
	middle := MiddleNum(data[:])
	assert.Equal(t, middle, float32(1.5))
}

func TestMiddleNum6(t *testing.T) {
	data := [...]int32{2, 2}
	middle := MiddleNum(data[:])
	assert.Equal(t, middle, float32(2))
}

func TestMiddleNum7(t *testing.T) {
	size := 33
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31() % 1000
	}
	fmt.Println(data)

	//
	middle := MiddleNum(data[:])

	// sort
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	// middle
	sortMiddle := data[size/2]
	fmt.Println(data)

	//
	assert.Equal(t, middle, float32(sortMiddle))
}

func TestMiddleNum8(t *testing.T) {
	size := 1024
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}

	//
	middle := MiddleNum(data[:])

	// sort
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	// middle
	sortMiddle := (float32(data[size/2-1]) + float32(data[size/2]))/2

	//
	assert.Equal(t, middle, float32(sortMiddle))
}
