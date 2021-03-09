package insertsort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestInsertSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	InsertSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestInsertSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	InsertSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestInsertSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	InsertSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestInsertSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	InsertSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestInsertSort5(t *testing.T) {
	size := 102400
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}
	InsertSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

