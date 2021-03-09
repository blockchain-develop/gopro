package poweringnumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowering_Naive1(t *testing.T) {
	fac := Powering_Naive(2, 0)
	assert.Equal(t, fac, 1)
}
func TestPowering_Naive2(t *testing.T) {
	fac := Powering_Naive(2, 1)
	assert.Equal(t, fac, 2)
}
func TestPowering_Naive3(t *testing.T) {
	fac := Powering_Naive(2, 10)
	assert.Equal(t, fac, 1024)
}

func TestPowering_DivideConquer1(t *testing.T) {
	fac := Powering_DivideConquer(2, 0)
	assert.Equal(t, fac, 1)
}
func TestPowering_DivideConquer2(t *testing.T) {
	fac := Powering_DivideConquer(2, 1)
	assert.Equal(t, fac, 2)
}
func TestPowering_DivideConquer3(t *testing.T) {
	fac := Powering_DivideConquer(2, 10)
	assert.Equal(t, fac, 1024)
}



