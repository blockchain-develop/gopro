package skiplist

import (
	"fmt"
	"math/rand"
)

func randLevel(max int) int {
	level := int(1)
	if (rand.Uint32() % 2 == 1) && (level <= max) {
		level ++
	}
	return level
}

type SkipListNode struct {
	key         int
	len         int
	forword     []*SkipListNode
}

func NewSkipListNode(key int, maxLevel int) *SkipListNode {
	level := randLevel(maxLevel)
	sln := &SkipListNode{
		key: key,
		forword: make([]*SkipListNode, level),
	}
	return sln
}

type SkipList struct {
	maxLevel     int
	len          int
	header       *SkipListNode
}

func (sl *SkipList) search(key int) []*SkipListNode {
	currentNode := sl.header
	updateNodes := make([]*SkipListNode, sl.maxLevel)
	for level := sl.maxLevel - 1;level >=0;level -- {
		for currentNode.forword[level] != nil && currentNode.forword[level].key < key {
			currentNode = currentNode.forword[level]
		}
		updateNodes[level] = currentNode
	}
	return updateNodes
}

func NewSkipList(maxLevel int) *SkipList {
	sl := &SkipList{
		maxLevel: maxLevel,
	}
	header := &SkipListNode{
		key: 0,
		forword: make([]*SkipListNode, maxLevel),
	}
	sl.header = header
	return sl
}

func (sl *SkipList) Insert(key int) bool {
	updateNodes := sl.search(key)
	if updateNodes[0].forword[0] != nil && updateNodes[0].forword[0].key == key {
		return false
	}
	sln := NewSkipListNode(key, sl.maxLevel)
	for i := 0;i < len(sln.forword);i ++ {
		sln.forword[i] = updateNodes[i].forword[i]
		updateNodes[i].forword[i] = sln
	}
	for i := len(sln.forword);i < sl.maxLevel;i ++ {
		updateNodes[i].len ++
	}
	sl.len ++
	return true
}

func (sl *SkipList) Delete(key int) bool {
	updateNodes := sl.search(key)
	if updateNodes[0].forword[0] == nil || updateNodes[0].forword[0].key != key {
		return false
	}
	currentNode := updateNodes[0].forword[0]
	for i := 0;i < len(currentNode.forword);i ++ {
		updateNodes[i].forword[i] = currentNode.forword[i]
		currentNode.forword[i] = nil
	}
	for i := len(currentNode.forword);i < sl.maxLevel;i ++ {
		updateNodes[i].len --
	}
	sl.len --
	return true
}

func (sl *SkipList) Exit(key int) bool {
	updateNodes := sl.search(key)
	if updateNodes[0].forword[0] == nil || updateNodes[0].forword[0].key != key {
		return false
	}
	return true
}

func (sl *SkipList) Print() {
	fmt.Printf("Skip list Len: %d \n [ ", sl.len)
	current := sl.header.forword[0]
	for current != nil {
		fmt.Printf("%d ", current.key)
		current = current.forword[0]
	}
	fmt.Printf("]\n")
}