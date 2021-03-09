package simpleselectsort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSimpleSelectSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	SimpleSelectSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestSimpleSelectSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	SimpleSelectSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestSimpleSelectSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	SimpleSelectSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestSimpleSelectSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	SimpleSelectSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestSimpleSelectSort5(t *testing.T) {
	size := 102400
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}
	SimpleSelectSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

