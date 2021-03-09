package mergesort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMergeSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	MergeSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestMergeSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	MergeSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestMergeSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	MergeSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestMergeSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	MergeSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestMergeSort5(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}
	MergeSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

