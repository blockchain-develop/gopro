package bucketsort

import (
	"github.com/gopro/memsort"
	"github.com/magiconair/properties/assert"
	"math/rand"
	"testing"
)

func TestBucketSort1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	BucketSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBucketSort2(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	BucketSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBucketSort3(t *testing.T) {
	data := [...]int32{5, 1, 3, 4, 2}
	BucketSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBucketSort4(t *testing.T) {
	data := [...]int32{5, 1, 3, 1, 2}
	BucketSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}

func TestBucketSort5(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}
	BucketSort(data[:])
	order := comsort.IsOrderly(data[:])
	assert.Equal(t, order, byte(1))
}


