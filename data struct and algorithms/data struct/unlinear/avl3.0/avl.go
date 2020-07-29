package avl3_0

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

type AVLNode struct {
	Value       int
	Data        string
	Left        *AVLNode
	Right       *AVLNode
}

func (node *AVLNode) Height() int {
	if node == nil {
		return 0
	}
	return max(node.Left.Height(), node.Right.Height()) + 1
}

func (node *AVLNode) Balance() int {
	return node.Left.Height() - node.Right.Height()
}

type AVLTree struct {
	Root       *AVLNode
}

func NewAVLTree() *AVLTree {
	tree := &AVLTree{
		Root: nil,
	}
	return tree
}

func (tree *AVLTree) Insert(value int, data string) bool {
	node, r := tree.insert(tree.Root, value, data)
	tree.Root = node
	return r
}

func (tree *AVLTree) insert(node *AVLNode, value int, data string) (*AVLNode, bool) {
	if node == nil {
		return &AVLNode{Value: value, Data: data,}, true
	}
	if value == node.Value {
		node.Data = data
		return node, false
	}
	var r bool
	if value < node.Value {
		node.Left, r = tree.insert(node.Left, value, data)
	} else {
		node.Right, r = tree.insert(node.Right, value, data)
	}
	node = tree.rebalance(node)
	return node, r
}

func (tree *AVLTree) rebalance(node *AVLNode) *AVLNode {
	switch {
	case node.Balance() == -2 && node.Right.Balance() == -1:
		return tree.rotateLeft(node)
	case node.Balance() == 2 && node.Left.Balance() == 1:
		return tree.rotateRight(node)
	case node.Balance() == -2 && node.Right.Balance() == 1:
		return tree.rotateRightLeft(node)
	case node.Balance() == 2 && node.Left.Balance() == -1:
		return tree.rotateLeftRight(node)
	}
	return node
}

func (tree *AVLTree) rotateLeft(node *AVLNode) *AVLNode {
	r := node.Right
	node.Right = r.Left
	r.Left = node
	return r
}

func (tree *AVLTree) rotateRight(node *AVLNode) *AVLNode {
	l := node.Left
	node.Left = l.Right
	l.Right = node
	return l
}

func (tree *AVLTree) rotateRightLeft(node *AVLNode) *AVLNode {
	node.Right = tree.rotateRight(node.Right)
	return tree.rotateLeft(node)
}

func (tree *AVLTree) rotateLeftRight(node *AVLNode) *AVLNode {
	node.Left = tree.rotateLeft(node.Left)
	return tree.rotateRight(node)
}

func (tree *AVLTree) Find(value int) *AVLNode {
	_, cur := tree.find(tree.Root, value)
	return cur
}

func (tree *AVLTree) find(node *AVLNode, value int) (*AVLNode, *AVLNode) {
	if node == nil {
		return nil, nil
	}
	if value == node.Value {
		return nil, node
	}
	var parent, cur *AVLNode
	if value < node.Value {
		parent, cur = tree.find(node.Left, value)
	} else {
		parent, cur = tree.find(node.Right, value)
	}
	if cur != nil && parent == nil {
		parent = node
	}
	return parent, cur
}

func (tree *AVLTree) next(node *AVLNode) (*AVLNode, *AVLNode) {
	if node == nil || node.Right == nil {
		return nil, nil
	}
	parent := node
	current := node.Right
	for current.Left != nil {
		parent = current
		current = current.Left
	}
	return parent, current
}
/*
This is not real pre node of node, just use for delete
 */
func (tree *AVLTree) pre(node *AVLNode) (*AVLNode, *AVLNode) {
	if node == nil || node.Left == nil {
		return nil, nil
	}
	parent := node
	current := node.Left
	for current.Right != nil {
		parent = current
		current = current.Right
	}
	return parent, current
}

func (tree *AVLTree) removeLeaf(leaf *AVLNode, parent *AVLNode) {
	if parent == nil {
		return
	}
	if leaf == parent.Left {
		parent.Left = nil
	} else if leaf == parent.Right{
		parent.Right = nil
	} else {

	}
}

func (tree *AVLTree) delete(node *AVLNode, value int) (*AVLNode, bool) {
	deleteNodeParent, deleteNode := tree.find(tree.Root, value)
	if deleteNode == nil {
		return node, false
	}
	replaceNodeParent, replaceNode := tree.next(deleteNode)
	if replaceNode == nil {
		replaceNodeParent, replaceNode = tree.pre(deleteNode)
	}
	if replaceNode == nil {
		if deleteNodeParent == nil {
			return nil, true
		}
		tree.removeLeaf(deleteNode, deleteNodeParent)
		return tree.fix(tree.Root, deleteNodeParent.Value), true
	}
	deleteNode.Value = replaceNode.Value
	deleteNode.Data = replaceNode.Data
	tree.removeLeaf(replaceNode, replaceNodeParent)
	return tree.fix(tree.Root, replaceNodeParent.Value), true
}

func (tree *AVLTree) fix(node *AVLNode, value int) (*AVLNode) {
	if node == nil {
		return node
	}
	if node.Value < value {
		node.Right = tree.fix(node.Right, value)
	} else if node.Value > value {
		node.Left = tree.fix(node.Left, value)
	}
	node = tree.rebalance(node)
	return node
}

func (tree *AVLTree) Delete(value int) bool {
	node, r := tree.delete(tree.Root, value)
	tree.Root = node
	return r
}