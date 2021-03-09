package maxminheap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func MaxHeapDump(mh *MaxHeap) []int {
	len := mh.Len()
	data := make([]int, len)
	for len > 0 {
		data[len - 1] = mh.Remove()
		len = mh.Len()
	}
	return data
}

func TestMaxHeap1(t *testing.T) {
	data := [...]int{1, 2, 3, 4, 5}
	mh := NewMaxHeap()
	mh.Init(data[:])
	r := MaxHeapDump(mh)
	for _, item := range r {
		fmt.Printf(" %d ", item)
	}
	fmt.Printf("\n")
	order := IsOrderly(r[:])
	assert.Equal(t, order, byte(1))
}

func TestMaxHeap2(t *testing.T) {
	data := [...]int{5, 4, 3, 2, 1}
	mh := NewMaxHeap()
	mh.Init(data[:])
	r := MaxHeapDump(mh)
	for _, item := range r {
		fmt.Printf(" %d ", item)
	}
	fmt.Printf("\n")
	order := IsOrderly(r[:])
	assert.Equal(t, order, byte(1))
}

func TestMaxHeap3(t *testing.T) {
	data := [...]int{5, 1, 3, 1, 2}
	mh := NewMaxHeap()
	mh.Init(data[:])
	r := MaxHeapDump(mh)
	for _, item := range r {
		fmt.Printf(" %d ", item)
	}
	fmt.Printf("\n")
	order := IsOrderly(r[:])
	assert.Equal(t, order, byte(1))
}

func TestMaxHeap4(t *testing.T) {
	size := 10240000
	data := make([]int, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int()
	}
	mh := NewMaxHeap()
	mh.Init(data[:])
	r := MaxHeapDump(mh)
	order := IsOrderly(r[:])
	assert.Equal(t, order, byte(1))
}

func TestMaxHeap5(t *testing.T) {
	mh := NewMaxHeap()
	mh.Insert(0)
	mh.Insert(5)
	mh.Insert(1)
	mh.Insert(3)
	mh.Insert(4)
	mh.Insert(2)
	mh.Insert(9)
	r := MaxHeapDump(mh)
	order := IsOrderly(r[:])
	assert.Equal(t, order, byte(1))
}

func TestMaxHeap6(t *testing.T) {
	size := 10240
	mh := NewMaxHeap()
	for i := 0;i < size;i ++ {
		mh.Insert(rand.Int())
	}
	r := MaxHeapDump(mh)
	order := IsOrderly(r[:])
	assert.Equal(t, order, byte(1))
}
