package shortestcommonsupersequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestCommonSupersequence_DP1(t *testing.T) {
	shortest := ShortestCommonSupersequence_DP([]byte("geek"), []byte("eke"))
	assert.Equal(t, shortest, 5)
}

func TestShortestCommonSupersequence_DP2(t *testing.T) {
	shortest := ShortestCommonSupersequence_DP([]byte("AGGTAB"), []byte("GXTXAYB"))
	assert.Equal(t, shortest, 9)
}