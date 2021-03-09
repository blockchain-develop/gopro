package maximumsubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumSubarray_Force1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	max, start, end := MaximumSubarray_Force(data[:])
	assert.Equal(t, max, 15)
	assert.Equal(t, start, 0)
	assert.Equal(t, end, 4)
}

func TestMaximumSubarray_Force2(t *testing.T) {
	data := [...]int32{-1, -2, 3, -4, -5}
	max, start, end := MaximumSubarray_Force(data[:])
	assert.Equal(t, max, 3)
	assert.Equal(t, start, 2)
	assert.Equal(t, end, 2)
}

func TestMaximumSubarray_Force3(t *testing.T) {
	data := [...]int32{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max, start, end := MaximumSubarray_Force(data[:])
	assert.Equal(t, max, 6)
	assert.Equal(t, start, 3)
	assert.Equal(t, end, 6)
}

func TestMaximumSubarray_Kadane1(t *testing.T) {
	data := [...]int{1, 2, 3, 4, 5}
	max := MaximumSubarray_Kadane(data[:])
	assert.Equal(t, max, 15)
}

func TestMaximumSubarray_Kadane2(t *testing.T) {
	data := [...]int{-1, -2, 3, -4, -5}
	max := MaximumSubarray_Kadane(data[:])
	assert.Equal(t, max, 3)
}

func TestMaximumSubarray_Kadane3(t *testing.T) {
	data := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max:= MaximumSubarray_Kadane(data[:])
	assert.Equal(t, max, 6)
}

func TestMaximumSubarray_KadaneWithIndex1(t *testing.T) {
	data := [...]int{1, 2, 3, 4, 5}
	max, start, end := MaximumSubarray_KadaneWithIndex(data[:])
	assert.Equal(t, max, 15)
	assert.Equal(t, start, 0)
	assert.Equal(t, end, 4)
}

func TestMaximumSubarray_KadaneWithIndex2(t *testing.T) {
	data := [...]int{-1, -2, 3, -4, -5}
	max, start, end := MaximumSubarray_KadaneWithIndex(data[:])
	assert.Equal(t, max, 3)
	assert.Equal(t, start, 2)
	assert.Equal(t, end, 2)
}

func TestMaximumSubarray_KadaneWithIndex3(t *testing.T) {
	data := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max, start, end := MaximumSubarray_KadaneWithIndex(data[:])
	assert.Equal(t, max, 6)
	assert.Equal(t, start, 3)
	assert.Equal(t, end, 6)
}
