package radixsort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestRadixSort1(t *testing.T) {
	data := [...]int32{11, 12, 13, 14, 15}
	RadixSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestRadixSort2(t *testing.T) {
	data := [...]int32{45, 14, 23, 22, 91}
	RadixSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestRadixSort3(t *testing.T) {
	data := [...]int32{55, 21, 63, 14, 72}
	RadixSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestRadixSort4(t *testing.T) {
	data := [...]int32{15, 31, 53, 31, 12}
	RadixSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestRadixSort5(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31() % 100000000
	}
	RadixSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}


