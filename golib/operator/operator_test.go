package operator

import (
	"fmt"
	"testing"
	"time"
)

func TestShift(t *testing.T) {
	n := 7235196223529507
	m := n >> 32
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

func TestNumberFromFigure(t *testing.T) {
	x := int64(1)
	for i := 0; i < 8; i++ {
		x *= 10
	}
	fmt.Printf("x: %d\n", x)
}
