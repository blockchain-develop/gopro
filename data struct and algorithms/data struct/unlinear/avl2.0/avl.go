package avl2_0

import "fmt"

const (
	LL_NOBALANCE = iota
	LR_NOBALANCE
	RR_NOBALANCE
	RL_NOBALANCE
	L_YESBALANCE
	R_YESBALANCE
	LR_NEWHEIGHT
)

type AVLNode struct {
	key      int
	left     *AVLNode
	right    *AVLNode
	balance  int
}

func NewAVLNode(key int) *AVLNode {
	node := &AVLNode{
		key: key,
		left: nil,
		right: nil,
		balance: 0,
	}
	return node
}

type AVLTree struct {
	root    *AVLNode
	size    int
}

func NewAVLTree() *AVLTree {
	tree := &AVLTree{
		root: nil,
		size: 0,
	}
	return tree
}

func (tree *AVLTree) typeAfterInsert(key int, node *AVLNode) int {
	if node.balance == 0 {
		return LR_NEWHEIGHT
	}
	if node.balance == 1 {
		if key > node.key {
			return R_YESBALANCE
		} else {
			if key < node.left.key {
				return LL_NOBALANCE
			} else {
				return LR_NOBALANCE
			}
		}
	} else {
		if key < node.key {
			return L_YESBALANCE
		} else {
			if key > node.right.key {
				return RR_NOBALANCE
			} else {
				return RL_NOBALANCE
			}
		}
	}
}

func (tree *AVLTree) rebalance(key int, node *AVLNode) {
	if node == nil || node.key == key {
		return
	}
	if key > node.key {
		if node.right == nil {
			return
		}
		node.balance = -1
		tree.rebalance(key, node.right)
	} else {
		if node.left == nil {
			return
		}
		node.balance = 1
		tree.rebalance(key, node.left)
	}
}

func (tree *AVLTree) rotateRight(node *AVLNode) *AVLNode {
	a,b := node, node.left
	a.left = b.right
	b.right = a
	return b
}

func (tree *AVLTree) rotateLeft(node *AVLNode) *AVLNode {
	a,b := node, node.right
	a.right = b.left
	b.left = a
	return b
}

func (tree *AVLTree) insert(key int, node *AVLNode) (*AVLNode, bool, bool) {
	if node == nil {
		return NewAVLNode(key), false, true
	}
	if key == node.key {
		return node, false, false
	}
	var rotated, result bool
	if key < node.key {
		node.left, rotated, result = tree.insert(key, node.left)
	} else {
		node.right, rotated, result = tree.insert(key, node.right)
	}
	if result == false {
		return node, rotated, result
	}
	if rotated == true {
		return node, rotated, result
	}
	if node.balance == 0 {
		if key > node.key {
			node.balance = -1
		} else if key < node.key {
			node.balance = 1
		}
		return node, rotated, result
	}
	afterInsert := tree.typeAfterInsert(key, node)
	switch  afterInsert {
	case L_YESBALANCE:
		node.balance = 0
	case R_YESBALANCE:
		node.balance = 0
	case LL_NOBALANCE:
		node = tree.rotateRight(node)
		node.balance = 0
		node.right.balance = 0
	case LR_NOBALANCE:
		b := node.left.right.balance
		node.left = tree.rotateLeft(node.left)
		node = tree.rotateRight(node)
		node.balance = 0
		if b == 0 {
			node.left.balance = 0
			node.right.balance = 0
		} else if b == 1 {
			node.left.balance = 0
			node.right.balance = -1
		} else {
			node.left.balance = 1
			node.right.balance = 0
		}
	case RR_NOBALANCE:
		node = tree.rotateLeft(node)
		node.balance = 0
		node.left.balance = 0
	case RL_NOBALANCE:
		b := node.right.left.balance
		node.right = tree.rotateRight(node.right)
		node = tree.rotateLeft(node)
		node.balance = 0
		if b == 0 {
			node.left.balance = 0
			node.right.balance = 0
		} else if b == 1 {
			node.left.balance = 0
			node.right.balance = -1
		} else {
			node.left.balance = 1
			node.right.balance = 0
		}
	}
	return node, true, result
}

func (tree *AVLTree) Insert(key int) bool {
	node, _, result:= tree.insert(key, tree.root)
	tree.root = node
	return result
}

func (tree *AVLTree) print(node *AVLNode) {
	if node == nil {
		return
	}
	fmt.Printf("{ %d(%d) ", node.key, node.balance)
	tree.print(node.left)
	tree.print(node.right)
	fmt.Printf(" }")
}

func (tree *AVLTree) Print() {
	tree.print(tree.root)
	fmt.Print("\n")
}
