package avl3_0

import (
	"testing"
)

func TestNewAVLTree1(t *testing.T) {
	avl := NewAVLTree()
	data := [...]int{3,2,1,4,5,6,7,10,9,8}
	for i := 0;i < len(data);i ++ {
		avl.Insert(data[i], "xxx")
	}
}

func TestNewAVLTree2(t *testing.T) {
	avl := NewAVLTree()
	data := [...]int{3,2,1,4,5,6,7,10,9,8}
	for i := 0;i < len(data);i ++ {
		avl.Insert(data[i], "xxx")
	}
	avl.Delete(7)
	avl.Delete(9)
	avl.Delete(10)
	avl.Delete(6)
	avl.Delete(4)
	avl.Delete(2)
	avl.Delete(1)
	avl.Delete(8)
	avl.Delete(5)
	avl.Delete(3)
}

