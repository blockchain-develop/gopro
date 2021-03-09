package longestpalindromicsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring_Naive1(t *testing.T) {
	data := []byte("hello")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 2)
	assert.Equal(t, length, 2)
}

func TestLongestPalindromicSubstring_Naive2(t *testing.T) {
	data := []byte("ABCDZJUDCBA")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 0)
	assert.Equal(t, length, 1)
}

func TestLongestPalindromicSubstring_Naive3(t *testing.T) {
	data := []byte("PATZJUJZTACCBCC")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 1)
	assert.Equal(t, length, 9)
}

func TestLongestPalindromicSubstring_Naive4(t *testing.T) {
	data := []byte("ABCDEEDCBA")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 0)
	assert.Equal(t, length, 10)
}

func TestLongestPalindromicSubstring_Naive5(t *testing.T) {
	data := []byte("ll")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 0)
	assert.Equal(t, length, 2)
}

func TestLongestPalindromicSubstring_Naive6(t *testing.T) {
	data := []byte("ab")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 0)
	assert.Equal(t, length, 1)
}

func TestLongestPalindromicSubstring_Naive7(t *testing.T) {
	data := []byte("a")
	index, length := LongestPalindromicSubstring_Naive(data[:])
	assert.Equal(t, index, 0)
	assert.Equal(t, length, 1)
}