package countingsort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCountingSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	CountingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestCountingSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	CountingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestCountingSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	CountingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestCountingSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	CountingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestCountingSort5(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31() % 10240000
	}
	CountingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}


