package xxxx

import (
	"fmt"
	"testing"
)

func TestXXXX(t *testing.T) {
	for i := 0;i < 10;i ++ {
		fmt.Printf("AAAAA")
		if i == 2 {
			fmt.Printf("BBBB")
			break
		}
		fmt.Printf("DDDD")
	}
	fmt.Printf("CCCC")
}
