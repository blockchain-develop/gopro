package btree

type BTreeNode struct {
	isLeaf          bool
	keys            []int
	childs          []*BTreeNode
}

func NewBTreeNode() *BTreeNode {
	node := &BTreeNode{
		isLeaf: true,
		keys: make([]int, 0),
		childs: make([]*BTreeNode, 0),
	}
	return node
}

type BTree struct {
	root            *BTreeNode
	t               int
}

func NewBTree() *BTree {
	tree := &BTree{
		root: NewBTreeNode(),
		t: 3,
	}
	return tree
}

func (bt *BTree) search(node *BTreeNode, key int) *BTreeNode {
	var index int
	var item int
	for index, item = range node.keys {
		if key <= item {
			break
		}
	}
	if item == key {
		return node
	}
	if node.isLeaf {
		return nil
	}
	if key > item {
		index += 1
	}
	return bt.search(node.childs[index], key)
}

func (bt *BTree) searchInsertLeaf(node *BTreeNode, key int) (*BTreeNode, int) {
	var index int
	var item int
	for index, item = range node.keys {
		if key <= item {
			break
		}
	}
	if item == key {
		return nil, 0
	}
	if key > item {
		index += 1
	}
	if node.isLeaf {
		return node, index
	}
	return bt.searchInsertLeaf(node.childs[index], key)
}

func (bt *BTree) Search(key int) bool {
	node := bt.search(bt.root, key)
	if node == nil {
		return false
	} else {
		return true
	}
}

func (bt *BTree) insert(node *BTreeNode, key int) (*BTreeNode, *BTreeNode, int) {
	var index int
	var item int
	for index, item = range node.keys {
		if key <= item {
			break
		}
	}
	if item == key {
		return nil, 0
	}
	if key > item {
		index += 1
	}
	if node.isLeaf {
		// insert
		node.keys = append(node.keys, 0)
		for i := len(node.keys) - 1;i > index;i -- {
			node.keys[i] = node.keys[i - 1]
		}
		node.keys[index] = key
		if len(node.keys) > bt.t * 2 {
			left := NewBTreeNode()
			right := NewBTreeNode()
			left.keys = append(left.keys, node.keys[0:bt.t]...)
			right.keys = append(right.keys, node.keys[bt.t + 1: bt.t]...)
			insertKey := node.keys[bt.t]
			return left, right, insertKey
		} else {
			return nil, nil, 0
		}
	}
	left, right, insertKey := bt.insert(node.childs[index], key)
	if left == nil || right == nil {
		return nil, nil, 0
	}
	// insert
	node.keys = append(node.keys, 0)
	for i := len(node.keys) - 1;i > index;i -- {
		node.keys[i] = node.keys[i - 1]
	}
	node.keys[index] = insertKey
	node.childs = append(node.childs, nil)
	for i := len(node.childs) - 1;i > index;i -- {
		node.keys[i] = node.keys[i - 1]
	}
	node.childs[index] = left
	node.childs[index +  1] = right
}

func (bt *BTree) Insert(key int) bool {
	leaf, index := bt.searchInsertLeaf(bt.root, key)
	if leaf == nil {
		return false
	}
}
