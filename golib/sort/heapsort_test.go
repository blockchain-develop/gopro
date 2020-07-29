package sort

import (
	"container/heap"
	"github.com/gopro/memsort"
	"github.com/magiconair/properties/assert"
	"math/rand"
	"testing"
)

type Int32Heap []int32
func (this Int32Heap) Len() int {
	return len(this)
}
func (this Int32Heap) Less(i, j int) bool {
	return this[i] < this[j]
}
func (this Int32Heap) Swap(i, j int) {
	this[i],this[j] = this[j], this[i]
}
func (this *Int32Heap) Push(x interface{}) {
	*this = append(*this, x.(int32))
}
func (this *Int32Heap) Pop() interface{} {
	old := *this
	n := len(old)
	x := old[n - 1]
	*this = old[0:n-1]
	return x
}

func TestHeapSort1(t *testing.T) {
	which := make(Int32Heap, 0)
	heap.Init(&which)

	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		heap.Push(&which, rand.Int31())
	}

	for i := 0;i < size;i ++ {
		x := heap.Pop(&which)
		data[i] = x.(int32)
	}
	order := comsort.IsOrderly(data)
	assert.Equal(t, order, byte(1))
}

func TestHeapSort2(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
	}

	which := Int32Heap(data)
	heap.Init(&which)

	data1 := make([]int32, size)
	for i := 0;i < size;i ++ {
		x := heap.Pop(&which)
		data1[i] = x.(int32)
	}
	order := comsort.IsOrderly(data1)
	assert.Equal(t, order, byte(1))
}
