package gcd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGCD1(t *testing.T) {
	gcd := GCD(3, 5)
	assert.Equal(t, gcd, 1)
}

func TestGCD2(t *testing.T) {
	gcd := GCD(5, 3)
	assert.Equal(t, gcd, 1)
}

func TestGCD3(t *testing.T) {
	gcd := GCD(3, 6)
	assert.Equal(t, gcd, 3)
}

func TestGCD4(t *testing.T) {
	gcd := GCD(6, 3)
	assert.Equal(t, gcd, 3)
}

func TestGCD5(t *testing.T) {
	gcd := GCD(105, 224)
	assert.Equal(t, gcd, 7)
}

