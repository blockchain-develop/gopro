package longestincreasingsubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIncreasingSubsequence_Recursive1(t *testing.T) {
	data := [...]int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	result := LongestIncreasingSubsequence_Recursive(data[:])
	assert.Equal(t, len(result), 6)
	assert.Equal(t, result[0], 80)
	assert.Equal(t, result[1], 60)
	assert.Equal(t, result[2], 41)
	assert.Equal(t, result[3], 33)
	assert.Equal(t, result[4], 22)
	assert.Equal(t, result[5], 10)
}
