package avl2_0

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
type AVLNode struct {
	key  uint32
	left *AVLNode
	right *AVLNode
}
func NewAVLNode(key uint32) *AVLNode {
	node := &AVLNode{
		key : key,
	}
	return node
}
type AVLTree struct {
	root *AVLNode
}
func NewAVLTree() *AVLTree {
	tree := &AVLTree{
		root : nil,
	}
	return tree
}
func (self *AVLTree) height(current *AVLNode) int {
	if current == nil {
		return 0
	}
	return max(self.height(current.left), self.height(current.right)) + 1
}
func (self *AVLTree) insert(current *AVLNode, item uint32) *AVLNode {
	if current == nil {
		node := NewAVLNode(item)
		return node
	}
	if item < current.key {
		current.left = self.insert(current.left, item)
	} else {
		current.right = self.insert(current.right, item)
	}
	if self.height(current.left) - self.height(current.right) == 2 {
		if current.right == nil {
			if item < current.left.key {
				newCurrrent := current.left
				newCurrrent.right = current
				current.left = nil
				current.right = nil
				return newCurrrent
			} else {
				newCurrent := current.left.right
				newCurrent.left = current.left
				newCurrent.right = current
				current.left.left = nil
				current.left.right = nil
				current.left = nil
				current.right = nil
				return newCurrent
			}
		} else {
			child := current.left
			lgrandchild := child.left
			rgrandchild := child.right
			if rgrandchild.left != nil {
				new_rgrandchild := rgrandchild.left
				new_rgrandchild.right = rgrandchild
				rgrandchild.left = nil
				rgrandchild = new_rgrandchild
				child.right = rgrandchild
			}
			if rgrandchild.right != nil {
				lgrandchild = child
				child = rgrandchild
				rgrandchild = child.right
				child.left = lgrandchild
				lgrandchild.right = nil
				current.left = child
			}
			newCurrent := current.left
			current.left = newCurrent.right
			newCurrent.right = current
			return newCurrent
		}
	} else if self.height(current.right) - self.height(current.left) == 2 {
		if current.left == nil {
			if item < current.right.key {
				newCurrent := current.right.left
				newCurrent.right = current.right
				newCurrent.left = current
				newCurrent.left.left = nil
				newCurrent.left.right = nil
				newCurrent.right.left = nil
				newCurrent.right.right = nil
				return newCurrent
			} else {
				newCurrent := current.right
				newCurrent.left = current
				current.right = nil
				return newCurrent
			}
		} else {
			child := current.right
			lgrandchild := child.left
			rgrandchild := child.right
			if lgrandchild.right != nil {
				new_lgrandchild := lgrandchild.right
				new_lgrandchild.left = lgrandchild
				lgrandchild.right = nil
				lgrandchild = new_lgrandchild
				child.left = lgrandchild
			}
			if lgrandchild.left != nil {
				rgrandchild = child
				child = lgrandchild
				lgrandchild = child.left
				child.right = rgrandchild
				rgrandchild.left = nil
				current.right = child
			}
			newCurrent := current.right
			current.right = newCurrent.left
			newCurrent.left = current
			return newCurrent
		}
	} else {
		return current
	}
}
func (self *AVLTree) Insert(item uint32) {
	self.root = self.insert(self.root, item)
}

func (self *AVLTree) deep_range(node *AVLNode) {
	if node == nil {
		return
	}
	self.deep_range(node.left)
	fmt.Printf("%d\t", node.key)
	self.deep_range(node.right)
}
func (self *AVLTree) Deep_Range() {
	self.deep_range(self.root)
}

func (self *AVLTree) degree_range(node *AVLNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d\t", node.key)
	self.degree_range(node.left)
	self.degree_range(node.right)
}
func (self *AVLTree) Degree_Range() {
	self.degree_range(self.root)
}

func (self *AVLTree) Degree_Range1() {
	q := queue.New()
	q.Enqueue(self.root)
	for q.Len() != 0 {
		node := q.Dequeue().(*AVLNode)
		if node == nil {
			continue
		}
		fmt.Printf("%d\t", node.key)
		q.Enqueue(node.left)
		q.Enqueue(node.right)
	}
}