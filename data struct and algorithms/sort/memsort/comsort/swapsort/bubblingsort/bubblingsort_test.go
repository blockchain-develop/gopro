package bubblingsort

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBubblingSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	BubblingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBubblingSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	BubblingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBubblingSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	BubblingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBubblingSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	BubblingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBubblingSort5(t *testing.T) {
	size := 102400
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}
	BubblingSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

