package longestcommonsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubstring_Naive1(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hello"), []byte("arm"))
	assert.Equal(t, si, -1)
	assert.Equal(t, ti, -1)
	assert.Equal(t, lenght, 0)
}
func TestLongestCommonSubstring_Naive2(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hello"), []byte("hello"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 5)
}
func TestLongestCommonSubstring_Naive3(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hello"), []byte("hel"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_Naive4(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hel"), []byte("hello"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_Naive5(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hello"), []byte("llo"))
	assert.Equal(t, si, 2)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_Naive6(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("llo"), []byte("hello"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 2)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_Naive7(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hello"), []byte("ll"))
	assert.Equal(t, si, 2)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 2)
}
func TestLongestCommonSubstring_Naive8(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("hello"), []byte("ablle"))
	assert.Equal(t, si, 2)
	assert.Equal(t, ti, 2)
	assert.Equal(t, lenght, 2)
}
func TestLongestCommonSubstring_Naive9(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("abcde"), []byte("edcba"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 4)
	assert.Equal(t, lenght, 1)
}
func TestLongestCommonSubstring_Naive10(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("abcdbcgh"), []byte("cbcbcgiu"))
	assert.Equal(t, si, 4)
	assert.Equal(t, ti, 3)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_Naive11(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_Naive([]byte("tutorialhorizon"), []byte("dynamictutorialProgramming"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 7)
	assert.Equal(t, lenght, 8)
}