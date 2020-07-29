package operator

import (
	"fmt"
	"testing"
	"time"
)

func TestShift(t *testing.T) {
	n := 4
	m := n << 1
	fmt.Printf("number %d left shift 1 is %d\n", n, m)
}

func TestShift1(t *testing.T) {
	n := 258
	a := n << 56
	a = a >> 56
	b := n & 0xFF
	fmt.Printf("number: %x, min bits: %x, min bits: %x\n", n, a, b)
}

func TestShift2(t *testing.T) {
	n := int(time.Now().Unix())
	n = n << 8
	m := 1000000
	m = m & 0xFF
	a := n + m
	fmt.Printf("%x %x %x\n", n, m, a)
}
