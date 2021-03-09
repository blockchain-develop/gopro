package longestpalindromicsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring_DP1(t *testing.T) {
	data := []byte("hello")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 2)
}

func TestLongestPalindromicSubstring_DP2(t *testing.T) {
	data := []byte("ABCDZJUDCBA")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 1)
}

func TestLongestPalindromicSubstring_DP3(t *testing.T) {
	data := []byte("PATZJUJZTACCBCC")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 9)
}

func TestLongestPalindromicSubstring_DP4(t *testing.T) {
	data := []byte("ABCDEEDCBA")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 10)
}

func TestLongestPalindromicSubstring_DP5(t *testing.T) {
	data := []byte("ll")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 2)
}

func TestLongestPalindromicSubstring_DP6(t *testing.T) {
	data := []byte("ab")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 1)
}

func TestLongestPalindromicSubstring_DP7(t *testing.T) {
	data := []byte("a")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 1)
}

func TestLongestPalindromicSubstring_DP8(t *testing.T) {
	data := []byte("baefeab")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 7)
}

func TestLongestPalindromicSubstring_DP9(t *testing.T) {
	data := []byte("forgeeksskeegfor")
	sub := LongestPalindromicSubstring_DP(data[:])
	assert.Equal(t, sub, 10)
}