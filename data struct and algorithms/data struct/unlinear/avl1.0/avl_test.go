package avl2_0

import (
	"fmt"
	"testing"
)

func TestNewAVLTree1(t *testing.T) {
	avl := NewAVLTree()
	data := [...]uint32{3,2,1,4,5,6,7,10,9,8}
	for i := 0;i < len(data);i ++ {
		avl.Insert(data[i])
	}
	fmt.Printf("Range av:\n")
	avl.Deep_Range()
	fmt.Printf("\n")
	fmt.Printf("Range av:\n")
	avl.Degree_Range()
	fmt.Printf("\n")
	fmt.Printf("Range av:\n")
	avl.Degree_Range1()
	fmt.Printf("\n")
}