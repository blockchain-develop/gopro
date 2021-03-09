package shortestcommonsupersequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestCommonSupersequence_Recursive1(t *testing.T) {
	shortest := ShortestCommonSupersequence_Recursive([]byte("geek"), []byte("eke"))
	assert.Equal(t, shortest, 5)
}

func TestShortestCommonSupersequence_Recursive2(t *testing.T) {
	shortest := ShortestCommonSupersequence_Recursive([]byte("AGGTAB"), []byte("GXTXAYB"))
	assert.Equal(t, shortest, 9)
}