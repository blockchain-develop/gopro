package longestcommonsubsequence

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence_Naive1(t *testing.T) {
	longest, subsequence := LongestCommonSubsequence_Naive([]byte("ABCBDAB"), []byte("BDCABA"))
	assert.Equal(t, longest, 4)
	// BDAB BCAB BCBA
	for _, item := range subsequence {
		fmt.Printf("%s\n", string(item))
	}
}

func TestLongestCommonSubsequence_Naive2(t *testing.T) {
	longest, subsequence := LongestCommonSubsequence_Naive([]byte("ABCDGH"), []byte("AEDFHR"))
	assert.Equal(t, longest, 3)
	// ADH
	for _, item := range subsequence {
		fmt.Printf("%s\n", string(item))
	}
}

func TestLongestCommonSubsequence_Naive3(t *testing.T) {
	longest, subsequence := LongestCommonSubsequence_Naive([]byte("GXTXAYB"), []byte("GTAB"))
	assert.Equal(t, longest, 4)
	// GTAB
	for _, item := range subsequence {
		fmt.Printf("%s\n", string(item))
	}
}