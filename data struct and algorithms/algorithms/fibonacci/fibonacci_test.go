package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci_Naive1(t *testing.T) {
	fac := Fibonacci_Naive(0)
	assert.Equal(t, fac, 0)
}

func TestFibonacci_Naive2(t *testing.T) {
	fac := Fibonacci_Naive(1)
	assert.Equal(t, fac, 1)
}

func TestFibonacci_Naive3(t *testing.T) {
	fac := Fibonacci_Naive(10)
	assert.Equal(t, fac, 55)
}

func TestFibonacci_Naive4(t *testing.T) {
	fac := Fibonacci_Naive(20)
	assert.Equal(t, fac, 6765)
}

func TestFibonacci_Tterative1(t *testing.T) {
	fac := Fibonacci_Tterative(0)
	assert.Equal(t, fac, 0)
}

func TestFibonacci_Tterative2(t *testing.T) {
	fac := Fibonacci_Tterative(1)
	assert.Equal(t, fac, 1)
}

func TestFibonacci_Tterative3(t *testing.T) {
	fac := Fibonacci_Tterative(10)
	assert.Equal(t, fac, 55)
}

func TestFibonacci_Tterative4(t *testing.T) {
	fac := Fibonacci_Tterative(20)
	assert.Equal(t, fac, 6765)
}

func TestFibonacci_Tterative5(t *testing.T) {
	fac := Fibonacci_Tterative(100)
	assert.Equal(t, fac, 3736710778780434371)
}
