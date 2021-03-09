package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFac_Tterative1(t *testing.T) {
	fac := Fac_Tterative(0)
	assert.Equal(t, fac, 1)
}

func TestFac_Tterative2(t *testing.T) {
	fac := Fac_Tterative(1)
	assert.Equal(t, fac, 1)
}

func TestFac_Tterative3(t *testing.T) {
	fac := Fac_Tterative(10)
	assert.Equal(t, fac, 3628800)
}

func TestFac_Recursive1(t *testing.T) {
	fac := Fac_Recursive(0)
	assert.Equal(t, fac, 1)
}

func TestFac_Recursive2(t *testing.T) {
	fac := Fac_Recursive(1)
	assert.Equal(t, fac, 1)
}

func TestFac_Recursive3(t *testing.T) {
	fac := Fac_Recursive(10)
	assert.Equal(t, fac, 3628800)
}
