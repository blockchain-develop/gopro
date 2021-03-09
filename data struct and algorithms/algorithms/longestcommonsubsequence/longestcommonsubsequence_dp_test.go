package longestcommonsubsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence_DP1(t *testing.T) {
	longest := LongestCommonSubsequence_DP([]byte("ABCBDAB"), []byte("BDCABA"))
	assert.Equal(t, longest, 4)
}

func TestLongestCommonSubsequence_DP2(t *testing.T) {
	longest := LongestCommonSubsequence_DP([]byte("ABCDGH"), []byte("AEDFHR"))
	assert.Equal(t, longest, 3)
}

func TestLongestCommonSubsequence_DP3(t *testing.T) {
	longest := LongestCommonSubsequence_DP([]byte("GXTXAYB"), []byte("GTAB"))
	assert.Equal(t, longest, 4)
}