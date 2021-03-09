package quicksort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	QuickSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestQuickSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	QuickSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestQuickSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	QuickSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestQuickSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	QuickSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestQuickSort5(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}
	QuickSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

