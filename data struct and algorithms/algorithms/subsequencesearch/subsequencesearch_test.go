package subsequencesearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsequenceSearch_Force1(t *testing.T) {
	source := []byte("hello")
	target := []byte("ll")
	index := SubsequenceSearch_Force(source, target)
	assert.Equal(t, index, 2)
}

func TestSubsequenceSearch_Force2(t *testing.T) {
	source := []byte("hello")
	target := []byte("hello")
	index := SubsequenceSearch_Force(source, target)
	assert.Equal(t, index, 0)
}

func TestSubsequenceSearch_Force3(t *testing.T) {
	source := []byte("hello")
	target := []byte("lo")
	index := SubsequenceSearch_Force(source, target)
	assert.Equal(t, index, 3)
}

func TestSubsequenceSearch_Force4(t *testing.T) {
	source := []byte("hello")
	target := []byte("oo")
	index := SubsequenceSearch_Force(source, target)
	assert.Equal(t, index, -1)
}

