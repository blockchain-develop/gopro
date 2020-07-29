package avl2_0

import (
	"testing"
)

func TestNewAVLTree1(t *testing.T) {
	avl := NewAVLTree()
	data := [...]int{3,2,1,4,5,6,7,10,9,8}
	for i := 0;i < len(data);i ++ {
		avl.Insert(data[i])
	}
	avl.Print()
}
