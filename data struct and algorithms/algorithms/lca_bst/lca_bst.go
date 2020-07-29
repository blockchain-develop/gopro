package lca_bst

type BSTNode struct {
	left      *BSTNode
	right     *BSTNode
	value     int
}

type BST struct {
	root      *BSTNode
}

func NewBST() *BST {
	bst := &BST{
		root: nil,
	}
	return bst
}

func (bst *BST) insert(value int, node *BSTNode) (*BSTNode, bool) {
	if node == nil {
		return &BSTNode{left:nil, right:nil, value: value}, true
	}
	var result bool
	if value < node.value {
		node.left, result = bst.insert(value, node.left)
	} else if value > node.value {
		node.right, result = bst.insert(value, node.right)
	} else {
		result = false
	}
	return node, result
}

func (bst *BST) Insert(value int) bool {
	node, result := bst.insert(value, bst.root)
	bst.root = node
	return result
}

func (bst *BST) find(value int, node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return node
	}
	if node.value < value {
		return bst.find(value, node.right)
	}
	if node.value > value {
		return bst.find(value, node.left)
	}
	return nil
}

func (bst *BST) Find(value int) *BSTNode {
	return bst.find(value, bst.root)
}

func (bst *BST) next(node *BSTNode) *BSTNode {
	if node == nil || node.right == nil {
		return nil
	}
	current := node.right
	for current.left != nil {
		current = current.left
	}
	return current
}

func (bst *BST) pre(node *BSTNode) *BSTNode {
	if node == nil || node.left == nil {
		return nil
	}
	current := node.left
	for current.right != nil {
		current = current.right
	}
	return current
}

func (bst *BST) lca(a int, b int, node *BSTNode) (int, bool) {
	if node == nil {
		return 0, false
	}
	if a < node.value && b < node.value {
		return bst.lca(a, b, node.left)
	} else if a > node.value && b > node.value {
		return bst.lca(a, b, node.right)
	} else {
		return node.value, true
	}
}

func (bst *BST) LCA(a int, b int) (int, bool) {
	return bst.lca(a, b, bst.root)
}

func LCA_BST(data []int, a int, b int) int {
	bst := NewBST()
	for _, v := range data {
		bst.Insert(v)
	}
	k, _ := bst.LCA(a, b)
	return k
}
