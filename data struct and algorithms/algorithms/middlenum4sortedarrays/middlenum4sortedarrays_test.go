package middlenum4sortedarrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleNum4SortedArrays1(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 3, 4, 5}
	middle_num := MiddleNum4SortedArrays(a[:], b[:])
	assert.Equal(t, middle_num, float32(3))
}

func TestMiddleNum4SortedArrays2(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{1, 2, 4, 5}
	middle_num := MiddleNum4SortedArrays(a[:], b[:])
	assert.Equal(t, middle_num, float32(3))
}

func TestMiddleNum4SortedArrays3(t *testing.T) {
	a := [...]int{1, 2, 4, 5}
	b := [...]int{1, 2, 4, 5}
	middle_num := MiddleNum4SortedArrays(a[:], b[:])
	assert.Equal(t, middle_num, float32(3))
}

func TestMiddleNum4SortedArrays4(t *testing.T) {
	a := [...]int{1, 2, 4, 5}
	b := [...]int{1}
	middle_num := MiddleNum4SortedArrays(a[:], b[:])
	assert.Equal(t, middle_num, float32(2))
}
