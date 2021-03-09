package knapsack01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsack01_DP1(t *testing.T) {
	w := [...]int{10, 20, 30}
	v := [...]int{60, 100, 120}
	c := 50
	result := Knapsack01_DP(w[:], v[:], c)
	assert.Equal(t, result, 220)
}

func TestKnapsack01_DP2(t *testing.T) {
	w := [...]int{4, 5, 7, 2, 8, 3, 9, 6, 1, 10}
	v := [...]int{25, 14, 15, 4, 14, 5, 14, 8, 1, 10}
	c := 34
	result := Knapsack01_DP(w[:], v[:], c)
	assert.Equal(t, result, 83)
}