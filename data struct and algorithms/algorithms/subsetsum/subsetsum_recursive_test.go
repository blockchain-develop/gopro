package subsetsum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetSum_Recursive1(t *testing.T) {
	data := [...]int{3, 34, 4, 12, 5, 2}
	result := SubsetSum_Recursive(data[:], 9)
	assert.Equal(t, result, true)
}

func TestSubsetSum_Recursive2(t *testing.T) {
	data := [...]int{3, 34, 4, 12, 5, 2}
	result := SubsetSum_Recursive(data[:], 10)
	assert.Equal(t, result, true)
}

func TestSubsetSum_Recursive3(t *testing.T) {
	data := [...]int{3, 34, 4, 12, 5, 2}
	result := SubsetSum_Recursive(data[:], 13)
	assert.Equal(t, result, false)
}