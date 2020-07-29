package rbtree1_0

type RBTreeNode struct {
	left      *RBTreeNode
	right     *RBTreeNode
	value     int
}

type RBTree struct {
	root      *RBTreeNode
}

func rotateLeft(x *RBTreeNode) *RBTreeNode {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func rotateRight(y *RBTreeNode) *RBTreeNode {
	x := y.left
	y.left = x.right
	x.right = y
	return x
}
