package middlenum

import "container/heap"

type Int32MinHeap []int32
func (h Int32MinHeap) Len() int {
	return len(h)
}
func (h Int32MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h Int32MinHeap) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *Int32MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int32))
}
func (h *Int32MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0:n-1]
	return x
}
func (h *Int32MinHeap) Fetch() int32 {
	if len(*h) == 0 {
		return 0
	}
	//x := (*h)[len(*h) - 1]
	x := (*h)[0]
	return x
}
type Int32MaxHeap []int32
func (h Int32MaxHeap) Len() int {
	return len(h)
}
func (h Int32MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h Int32MaxHeap) Swap(i, j int) {
	h[i],h[j] = h[j],h[i]
}
func (h *Int32MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int32))
}
func (h *Int32MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0:n-1]
	return x
}
func (h *Int32MaxHeap) Fetch() int32 {
	if len(*h) == 0 {
		return 0
	}
	//x := (*h)[len(*h) - 1]
	x := (*h)[0]
	return x
}

func MiddleNum(data []int32) float32 {
	maxHeap := make(Int32MaxHeap, 0)
	heap.Init(&maxHeap)
	minHeap := make(Int32MinHeap, 0)
	heap.Init(&minHeap)

	maxInMaxHeap := int32(0)
	minInMinHeap := int32(0)
	for _, item := range data {
		if maxHeap.Len() == minHeap.Len() {
			if item <= maxInMaxHeap {
				heap.Push(&maxHeap, item)
				maxInMaxHeap = maxHeap.Fetch()
			} else {
				heap.Push(&minHeap, item)
				minInMinHeap = minHeap.Fetch()
			}
		} else if maxHeap.Len() > minHeap.Len() {
			if item < maxInMaxHeap {
				which := heap.Pop(&maxHeap)
				heap.Push(&minHeap, which)
				heap.Push(&maxHeap, item)
				maxInMaxHeap = maxHeap.Fetch()
				minInMinHeap = minHeap.Fetch()
			} else {
				heap.Push(&minHeap, item)
				minInMinHeap = minHeap.Fetch()
			}
		} else {
			if item > minInMinHeap {
				which := heap.Pop(&minHeap)
				heap.Push(&maxHeap, which)
				heap.Push(&minHeap, item)
				maxInMaxHeap = maxHeap.Fetch()
				minInMinHeap = minHeap.Fetch()
			} else {
				heap.Push(&maxHeap, item)
				maxInMaxHeap = maxHeap.Fetch()
			}
		}
	}

	if maxHeap.Len() == minHeap.Len() {
		return (float32(minHeap.Fetch()) + float32(maxHeap.Fetch()))/ 2
	} else if maxHeap.Len() > minHeap.Len() {
		return float32(maxHeap.Fetch())
	} else {
		return float32(minHeap.Fetch())
	}
}