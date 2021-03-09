package longestcommonsubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubstring_DP1(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hello"), []byte("arm"))
	assert.Equal(t, si, -1)
	assert.Equal(t, ti, -1)
	assert.Equal(t, lenght, 0)
}
func TestLongestCommonSubstring_DP2(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hello"), []byte("hello"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 5)
}
func TestLongestCommonSubstring_DP3(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hello"), []byte("hel"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_DP4(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hel"), []byte("hello"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_DP5(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hello"), []byte("llo"))
	assert.Equal(t, si, 2)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_DP6(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("llo"), []byte("hello"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 2)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_DP7(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hello"), []byte("ll"))
	assert.Equal(t, si, 2)
	assert.Equal(t, ti, 0)
	assert.Equal(t, lenght, 2)
}
func TestLongestCommonSubstring_DP8(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("hello"), []byte("ablle"))
	assert.Equal(t, si, 2)
	assert.Equal(t, ti, 2)
	assert.Equal(t, lenght, 2)
}
func TestLongestCommonSubstring_DP9(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("abcde"), []byte("edcba"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 4)
	assert.Equal(t, lenght, 1)
}
func TestLongestCommonSubstring_DP10(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("abcdbcgh"), []byte("cbcbcgiu"))
	assert.Equal(t, si, 4)
	assert.Equal(t, ti, 3)
	assert.Equal(t, lenght, 3)
}
func TestLongestCommonSubstring_DP11(t *testing.T) {
	si, ti, lenght := LongestCommonSubstring_DP([]byte("tutorialhorizon"), []byte("dynamictutorialProgramming"))
	assert.Equal(t, si, 0)
	assert.Equal(t, ti, 7)
	assert.Equal(t, lenght, 8)
}