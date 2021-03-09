package knapsackn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsackN_DP1(t *testing.T) {
	w := [...]int{10, 20, 30}
	v := [...]int{60, 100, 120}
	n := [...]int{1, 1, 1}
	c := 50
	result := KnapsackN_DP(w[:], v[:], n[:], c)
	assert.Equal(t, result, 220)
}

func TestKnapsackN_DP2(t *testing.T) {
	w := [...]int{4, 5, 7, 2, 8, 3, 9, 6, 1, 10}
	v := [...]int{25, 14, 15, 4, 14, 5, 14, 8, 1, 10}
	n := [...]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	c := 34
	result := KnapsackN_DP(w[:], v[:], n[:], c)
	assert.Equal(t, result, 83)
}

func TestKnapsackN_DP3(t *testing.T) {
	w := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	v := [...]int{1, 5, 8, 9, 10, 17, 17, 20}
	n := [...]int{100, 100, 100, 100, 100, 100, 100, 100}
	c := 8
	result := KnapsackN_DP(w[:], v[:], n[:], c)
	assert.Equal(t, result, 22)
}

func TestKnapsackN_DP4(t *testing.T) {
	w := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	v := [...]int{3, 5, 8, 9, 10, 17, 17, 20}
	n := [...]int{100, 100, 100, 100, 100, 100, 100, 100}
	c := 8
	result := KnapsackN_DP(w[:], v[:], n[:], c)
	assert.Equal(t, result, 24)
}