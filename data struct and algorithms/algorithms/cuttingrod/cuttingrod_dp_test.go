package cuttingrod

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCuttingRod_DP1(t *testing.T) {
	l := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	v := [...]int{1, 5, 8, 9, 10, 17, 17, 20}
	rodl := 8
	max_value := CuttingRod_DP(l[:], v[:], rodl)
	assert.Equal(t, max_value, 22)
}

func TestCuttingRod_DP2(t *testing.T) {
	l := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	v := [...]int{3, 5, 8, 9, 10, 17, 17, 20}
	rodl := 8
	max_value := CuttingRod_DP(l[:], v[:], rodl)
	assert.Equal(t, max_value, 24)
}