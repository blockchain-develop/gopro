package print

import (
	"fmt"
	"testing"
)

func TestPrintFormat(t *testing.T) {
	n := 12
	fmt.Printf("%d\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%c\n", n)
	fmt.Printf("%x\n", n)
}
