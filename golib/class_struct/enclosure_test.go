package class_struct

import (
	"fmt"
	"testing"
)

type AAA struct {
	a string
}

func (a *AAA) print() {
	fmt.Println(a.a)
}

func (a *AAA) print1() {
	fmt.Println("other", a.a)
}

type BBB struct {
	AAA
	b string
}

func (b *BBB) print() {
	fmt.Println(b.b)
}

func TestEnclosure(t *testing.T) {
	b := &BBB{
		AAA: AAA{"aaa"},
		b:   "bbb",
	}
	b.print()
	b.print1()
}