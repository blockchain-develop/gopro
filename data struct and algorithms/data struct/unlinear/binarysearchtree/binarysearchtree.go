package binarysearchtree

type BinarySearchTreeNode struct {
	key int32
	left *BinarySearchTreeNode
	right *BinarySearchTreeNode
	parent *BinarySearchTreeNode
}
func NewNode(key int32) *BinarySearchTreeNode {
	node := &BinarySearchTreeNode{
		key : key,
		left : nil,
		right : nil,
		parent: nil,
	}
	return node
}
type BinarySearchTree struct {
	root *BinarySearchTreeNode
}
func NewBinarySearchTree() *BinarySearchTree {
	tree := &BinarySearchTree {
		root : nil,
	}
	return tree
}
func (self *BinarySearchTree) InSert(key int32) {
	node := NewNode(key)
	if self.root == nil {
		self.root = node
		return
	}
	current := self.root
	for true {
		if key > current.key {
			if current.right == nil {
				current.right = node
				return
			} else {
				current = current.right
			}
		} else if key < current.key {
			if current.left == nil {
				current.left = node
				return
			} else {
				current = current.left
			}
		} else {
			return
		}
	}
}
func (self *BinarySearchTree) Find(key int32) *BinarySearchTreeNode {
	current := self.root
	for current != nil {
		if key > current.key {
			current = current.right
		} else if key < current.key {
			current = current.left
		} else {
			break
		}
	}
	return current
}
func (self *BinarySearchTree) Delete(key int32) *BinarySearchTreeNode {
	return nil
}
func (self *BinarySearchTree) FindMin(x *BinarySearchTreeNode) *BinarySearchTreeNode {
	if x == nil {
		return nil
	}
	current := x
	for current.left != nil {
		current = current.left
	}
	return current
}
func (self *BinarySearchTree) DeleteMin() {

}
func (self *BinarySearchTree) NextLarger(x *BinarySearchTreeNode) *BinarySearchTreeNode {
	if x == nil {
		return nil
	}
	if x.right != nil {
		return self.FindMin(x.right)
	}
	y := x.parent
	for y != nil && y.left != x {
		x = y
		y = x.parent
	}
	return y
}
func (self *BinarySearchTree) Get() []int32 {
	data := make([]int32, 0)
	self.get(self.root, data)
	return data
}
func (self *BinarySearchTree) get(node *BinarySearchTreeNode, data []int32) {
	if node.left != nil {
		self.get(node.left, data)
	}
	data = append(data, node.key)
	if node.right != nil {
		self.get(node.right, data)
	}
}
