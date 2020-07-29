package avl4_0

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type AVLNode struct {
	key     int
	height  int
	left    *AVLNode
	right   *AVLNode
}

func NewAVLNode(key int) *AVLNode {
	node := &AVLNode{
		key: key,
		height: 0,
	}
	return node
}

type AVLTree struct {
	root  *AVLNode
}

func NewAVLTree() *AVLTree {
	tree := &AVLTree{
	}
	return tree
}

func (tree *AVLTree) height(node *AVLNode) int {
	if node == nil {
		return -1
	} else {
		return node.height
	}
}

func (tree *AVLTree) calculateHeight(node *AVLNode) {
	node.height = max(tree.height(node.left), tree.height(node.right)) + 1
}

func (tree *AVLTree) balance(node *AVLNode) int {
	return tree.height(node.left) - tree.height(node.right)
}

func (tree *AVLTree) rotate_right(node *AVLNode) *AVLNode {
	root := node.left
	node.left = root.right
	root.right = node
	tree.calculateHeight(node)
	tree.calculateHeight(root)
	return root
}

func (tree *AVLTree) rotate_left(node *AVLNode) *AVLNode {
	root := node.right
	node.right = root.left
	root.left = node
	tree.calculateHeight(node)
	tree.calculateHeight(root)
	return root
}

func (tree *AVLTree) insert(node *AVLNode, key int) *AVLNode {
	if node == nil {
		node := &AVLNode{
			key: key,
			height: 0,
		}
		return node
	}
	if key < node.key {
		node.left = tree.insert(node.left, key)
	} else {
		node.right = tree.insert(node.right, key)
	}
	tree.calculateHeight(node)
	balance := tree.balance(node)
	if balance == 2 {
		if tree.balance(node.left) == 1 {
			// LL case
			node = tree.rotate_right(node)
		} else {
			// LR case
			node.left = tree.rotate_left(node.left)
			node = tree.rotate_right(node)

		}
	} else if balance == -2 {
		if tree.balance(node.right) == 1 {
			// RL case
			node.right = tree.rotate_right(node.right)
			node = tree.rotate_left(node)
		} else {
			// RR case
			node = tree.rotate_left(node)
		}
	}
	return node
}

func (tree *AVLTree) Insert(key int) {
	if tree.root == nil {
		tree.root = NewAVLNode(key)
		return
	}
	tree.root = tree.insert(tree.root, key)
}
