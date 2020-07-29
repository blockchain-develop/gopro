package minimumspanningtrees

import (
	"container/heap"
	"math"
)

type QueueItem struct {
	key      *SpanNode
}

type Queue []*QueueItem

func (q *Queue) Add(key *SpanNode) {
	item := &QueueItem{
		key: key,
	}
	*q = append(*q, item)
}

func (q *Queue) Remove() *SpanNode {
	if len(*q) == 0 {
		return nil
	}
	key := (*q)[0].key
	*q = (*q)[1:]
	return key
}

type SpanNode struct {
	index     int
	child     []*SpanNode
}

func (sn *SpanNode) DepthFirstSearch(data *[]int) {
	if sn == nil {
		return
	}
	*data = append(*data, sn.index)
	for _, item := range sn.child {
		item.DepthFirstSearch(data)
	}
}

func (sn *SpanNode) DegreeFirstSearch(data *[]int) {
	q := make(Queue, 0)
	q.Add(sn)
	node := q.Remove()
	for node != nil {
		*data = append(*data, node.index)
		for _, child := range node.child {
			q.Add(child)
		}
		node = q.Remove()
	}
}

type PriorityItem struct {
	key       int
	index     int
	a         int
	b         int
}

type PriorityQueue []*PriorityItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].key <= pq[j].key
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityItem)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	len := len(*pq)
	if len == 0 {
		return nil
	}
	item := (*pq)[len - 1]
	*pq = (*pq)[0:len-1]
	return item
}

func (pq *PriorityQueue) Update(x interface{}, key int) {
	item := x.(*PriorityItem)
	item.key = key
	heap.Fix(pq, item.index)
}

func MinimumSpanningTrees_Prim(graph [][]int) *SpanNode {
	v := len(graph)
	setQ := make([]*PriorityItem, v)
	setA := make([]*PriorityItem, 0)
	pq := make(PriorityQueue, v)
	for i := 0;i < v;i ++ {
		item := &PriorityItem{
			key : math.MaxInt32,
			index: i,
			a: i,
			b: -1,
		}
		setQ[i] = item
		pq[i] = item
	}
	pq[0].key = 0
	heap.Init(&pq)

	//
	len := pq.Len()
	for len > 0 {
		x := heap.Pop(&pq)
		for i,k := range graph[x.(*PriorityItem).a] {
			if setQ[i] != nil && k > 0 && setQ[i].key > k {
				item := setQ[i]
				item.key = k
				item.b = x.(*PriorityItem).a
				heap.Fix(&pq, item.index)
			}
		}
		setQ[x.(*PriorityItem).a] = nil
		setA = append(setA, x.(*PriorityItem))
		len = pq.Len()
	}

	// build span tree
	var root *SpanNode
	nodes := make([]*SpanNode, v)
	for _, item := range setA {
		node := &SpanNode{
			index: item.a,
		}
		if item.b == -1 {
			root = node
		} else {
			nodes[item.b].child = append(nodes[item.b].child, node)
		}
		nodes[item.a] = node
	}
	return root
}
