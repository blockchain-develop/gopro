package longestincreasingsubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestIncreasingSubsequence_Recursive2(t *testing.T) {
	data := [...]int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	result := LongestIncreasingSubsequence_Recursive2(data[:])
	assert.Equal(t, result, 6)
}
