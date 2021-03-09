package binarysearchtree

import (
	"github.com/gopro/memsort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearchTree1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	tree := NewBinarySearchTree()
	for i := 0;i < len(data);i ++ {
		tree.InSert(data[i])
	}
	{
		node := tree.Find(1)
		assert.Equal(t, node.key, int32(1))
	}
	{
		node := tree.Find(3)
		assert.Equal(t, node.key, int32(3))
	}
	{
		node := tree.Find(5)
		assert.Equal(t, node.key, int32(5))
	}
	{
		node := tree.FindMin(tree.root)
		assert.Equal(t, node.key, int32(1))
	}
}

func TestBinarySearchTree2(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	tree := NewBinarySearchTree()
	for i := 0;i < len(data);i ++ {
		tree.InSert(data[i])
	}
	sort_data := tree.Get()
	order := comsort.IsOrderly(sort_data)
	assert.Equal(t, order, byte(1))
}

func TestBinarySearchTree3(t *testing.T) {
	data := [...]int32{9, 2, 3, 4, 5}
	tree := NewBinarySearchTree()
	for i := 0;i < len(data);i ++ {
		tree.InSert(data[i])
	}
	sort_data := tree.Get()
	order := comsort.IsOrderly(sort_data)
	assert.Equal(t, order, byte(1))
}

func TestBinarySearchTree4(t *testing.T) {
	data := [...]int32{5, 4, 3, 2, 1}
	tree := NewBinarySearchTree()
	for i := 0;i < len(data);i ++ {
		tree.InSert(data[i])
	}
	sort_data := tree.Get()
	order := comsort.IsOrderly(sort_data)
	assert.Equal(t, order, byte(1))
}
