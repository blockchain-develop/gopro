package shortestcommonsupersequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestCommonSupersequence_L1(t *testing.T) {
	shortest := ShortestCommonSupersequence_L([]byte("geek"), []byte("eke"))
	assert.Equal(t, shortest, 5)
}

func TestShortestCommonSupersequence_L2(t *testing.T) {
	shortest := ShortestCommonSupersequence_L([]byte("AGGTAB"), []byte("GXTXAYB"))
	assert.Equal(t, shortest, 9)
}