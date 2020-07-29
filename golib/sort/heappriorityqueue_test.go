package sort

import (
	"container/heap"
	"github.com/gopro/memsort"
	"github.com/magiconair/properties/assert"
	"math/rand"
	"testing"
)

type Item struct {
	value int32
	priority int
	index int
}
type PriorityQueue []*Item
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i],pq[j] = pq[j],pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n - 1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n - 1]
	return item
}
func (pq *PriorityQueue) Update(item *Item, value int32, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func TestHeapPriorityQueue1(t *testing.T) {
	size := 102400
	pq := make(PriorityQueue, size)
	for i := 0;i < size;i ++ {
		pq[i] = &Item {
			value: rand.Int31(),
			priority: 2,
			index: i,
		}
	}

	heap.Init(&pq)

	item := &Item {
		value: 1,
		priority: 1,
	}
	heap.Push(&pq, item)

	pq.Update(item, item.value, 3)

	data1 := make([]int32, size + 1)
	for i := 0;i < size + 1;i ++ {
		x := heap.Pop(&pq)
		data1[i] = int32((x.(*Item).priority))
	}
	order := comsort.IsOrderly(data1)
	assert.Equal(t, order, byte(2))
}
