package coinchange

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsackN_DP1(t *testing.T) {
	coins := [...]int{1, 2, 3}
	amount := 4
	result := CoinChange_DP(coins[:], amount)
	assert.Equal(t, result, 4)
}

func TestKnapsackN_DP2(t *testing.T) {
	coins := [...]int{2, 5, 3, 6}
	amount := 10
	result := CoinChange_DP(coins[:], amount)
	assert.Equal(t, result, 5)
}
