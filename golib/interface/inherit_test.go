package _interface

import (
	"fmt"
	"testing"
)

type ImmutableTree struct {
	Root        int
	Version     int
}

func (t *ImmutableTree) clone() *ImmutableTree {
	return &ImmutableTree{
		Root:    t.Root,
		Version: t.Version,
	}
}

type MutableTree struct {
	*ImmutableTree
	lastSaved     *ImmutableTree
}

func TestInherit(t *testing.T) {
	immutableTree := &ImmutableTree {
		Root: 0,
		Version: 0,
	}
	mutableTree := &MutableTree{
		ImmutableTree: immutableTree,
		lastSaved: immutableTree.clone(),
	}
	fmt.Printf("root: %d, version: %d\n", mutableTree.Root, mutableTree.Version)
	fmt.Printf("root: %d, version: %d\n", mutableTree.ImmutableTree.Root, mutableTree.ImmutableTree.Version)
	fmt.Printf("root: %d, version: %d\n", mutableTree.lastSaved.Root, mutableTree.lastSaved.Version)

	mutableTree.ImmutableTree.Root ++
	mutableTree.ImmutableTree.Version ++
	fmt.Printf("root: %d, version: %d\n", mutableTree.Root, mutableTree.Version)
	fmt.Printf("root: %d, version: %d\n", mutableTree.ImmutableTree.Root, mutableTree.ImmutableTree.Version)
	fmt.Printf("root: %d, version: %d\n", mutableTree.lastSaved.Root, mutableTree.lastSaved.Version)
}
