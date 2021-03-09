package subsetsum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetSum_DP1(t *testing.T) {
	data := [...]int{3, 34, 4, 12, 5, 2}
	result := SubsetSum_DP(data[:], 9)
	assert.Equal(t, result, true)
}

func TestSubsetSum_DP2(t *testing.T) {
	data := [...]int{3, 34, 4, 12, 5, 2}
	result := SubsetSum_DP(data[:], 10)
	assert.Equal(t, result, true)
}

func TestSubsetSum_DP3(t *testing.T) {
	data := [...]int{3, 34, 4, 12, 5, 2}
	result := SubsetSum_DP(data[:], 13)
	assert.Equal(t, result, false)
}