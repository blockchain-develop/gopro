package minimumnumbercoins

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumNumberCoins1(t *testing.T) {
	value := 40
	vrs := []int{1, 10, 40}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{40})
}

func TestMinimumNumberCoins2(t *testing.T) {
	value := 40
	vrs := []int{1, 10, 20}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{20, 20})
}

func TestMinimumNumberCoins3(t *testing.T) {
	value := 40
	vrs := []int{1, 10, 20, 80}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{20, 20})
}

func TestMinimumNumberCoins4(t *testing.T) {
	value := 40
	vrs := []int{1, 10, 20, 30}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{30, 10})
}

func TestMinimumNumberCoins5(t *testing.T) {
	value := 51
	vrs := []int{1, 10, 20, 30}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{30, 20, 1})
}

func TestMinimumNumberCoins6(t *testing.T) {
	value := 52
	vrs := []int{1, 10, 20, 30}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{30, 20, 1, 1})
}

func TestMinimumNumberCoins7(t *testing.T) {
	value := 162
	vrs := []int{1, 10, 20, 30}
	coins := MinimumNumberCoins(value, vrs)
	assert.Equal(t, coins, []int{30, 30, 30, 30, 30, 10, 1, 1})
}




