package longestpalindromicsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring_Force1(t *testing.T) {
	data := []byte("hello")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("ll"))
}

func TestLongestPalindromicSubstring_Force2(t *testing.T) {
	data := []byte("ABCDZJUDCBA")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("A"))
}

func TestLongestPalindromicSubstring_Force3(t *testing.T) {
	data := []byte("PATZJUJZTACCBCC")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("ATZJUJZTA"))
}

func TestLongestPalindromicSubstring_Force4(t *testing.T) {
	data := []byte("ABCDEEDCBA")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("ABCDEEDCBA"))
}

func TestLongestPalindromicSubstring_Force5(t *testing.T) {
	data := []byte("ll")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("ll"))
}

func TestLongestPalindromicSubstring_Force6(t *testing.T) {
	data := []byte("ab")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("a"))
}

func TestLongestPalindromicSubstring_Force7(t *testing.T) {
	data := []byte("a")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("a"))
}

func TestLongestPalindromicSubstring_Force8(t *testing.T) {
	data := []byte("baefeab")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("baefeab"))
}

func TestLongestPalindromicSubstring_Force9(t *testing.T) {
	data := []byte("forgeeksskeegfor")
	sub := LongestPalindromicSubstring_Force(data[:])
	assert.Equal(t, sub, []byte("geeksskeeg"))
}