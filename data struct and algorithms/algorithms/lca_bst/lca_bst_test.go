package lca_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCA_BST1(t *testing.T) {
	data := [...]int{20,22,8,4,12,10,14}
	k := LCA_BST(data[:],10, 14)
	assert.Equal(t, k, 12)
}

func TestLCA_BST2(t *testing.T) {
	data := [...]int{20,22,8,4,12,10,14}
	k := LCA_BST(data[:],14, 8)
	assert.Equal(t, k, 8)
}

func TestLCA_BST3(t *testing.T) {
	data := [...]int{20,22,8,4,12,10,14}
	k := LCA_BST(data[:],10, 22)
	assert.Equal(t, k, 20)
}
