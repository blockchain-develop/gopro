package longestincreasingsubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIncreasingSubsequence_Naive1(t *testing.T) {
	data := [...]int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	result := LongestIncreasingSubsequence_Naive(data[:])
	assert.Equal(t, len(result), 6)
	assert.Equal(t, result[0], 10)
	assert.Equal(t, result[1], 22)
	assert.Equal(t, result[2], 33)
	assert.Equal(t, result[3], 50)
	assert.Equal(t, result[4], 60)
	assert.Equal(t, result[5], 80)
}
