package editdistance

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence_Recursive1(t *testing.T) {
	longest, subsequence := LongestCommonSubsequence_Recursive([]byte("ABCBDAB"), []byte("BDCABA"))
	assert.Equal(t, longest, 4)
	// BDAB BCAB BCBA
	for _, item := range subsequence {
		fmt.Printf("%s\n", string(item))
	}
}

func TestLongestCommonSubsequence_Recursive2(t *testing.T) {
	longest, subsequence := LongestCommonSubsequence_Recursive([]byte("ABCDGH"), []byte("AEDFHR"))
	assert.Equal(t, longest, 3)
	// ADH
	for _, item := range subsequence {
		fmt.Printf("%s\n", string(item))
	}
}

func TestLongestCommonSubsequence_Recursive3(t *testing.T) {
	longest, subsequence := LongestCommonSubsequence_Recursive([]byte("GXTXAYB"), []byte("GTAB"))
	assert.Equal(t, longest, 4)
	// GTAB
	for _, item := range subsequence {
		fmt.Printf("%s\n", string(item))
	}
}