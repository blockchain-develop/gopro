package editdistance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditDistance_DP1(t *testing.T) {
	longest := EditDistance_DP([]byte("geek"), []byte("gesek"))
	assert.Equal(t, longest, 1)
}

func TestEditDistance_DP2(t *testing.T) {
	longest := EditDistance_DP([]byte("cat"), []byte("cut"))
	assert.Equal(t, longest, 1)
}

func TestEditDistance_DP3(t *testing.T) {
	longest := EditDistance_DP([]byte("sunday"), []byte("saturday"))
	assert.Equal(t, longest, 3)
}