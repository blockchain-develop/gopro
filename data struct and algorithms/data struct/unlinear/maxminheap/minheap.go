package maxminheap

type MinHeap struct {
	data       []int
}

func NewMinHeap() *MinHeap {
	mh := &MinHeap{
		data: make([]int, 0),
	}
	return mh
}

func (mh *MinHeap) fix(parent int) bool {
	child := parent * 2 + 1
	updated := false
	for child < len(mh.data) {
		if len(mh.data) > child+1 && mh.data[child+1] < mh.data[child] {
			child = child + 1
		}
		if mh.data[parent] < mh.data[child] {
			return updated
		}
		mh.data[parent], mh.data[child] = mh.data[child], mh.data[parent]
		parent = child
		child = parent*2 + 1
		updated = true
	}
	return updated
}

func (mh *MinHeap) Init(data []int) {
	mh.data = append(mh.data, data...)
	for i := len(mh.data) / 2 - 1;i >= 0;i -- {
		mh.fix(i)
	}
}

func (mh *MinHeap) Insert(item int) {
	mh.data = append(mh.data, item)
	pos := len(mh.data) / 2 - 1
	updated := true
	for pos >= 0 && updated == true {
		updated = mh.fix(pos)
		pos = (pos + 1) / 2 - 1
	}
}

func (mh *MinHeap) Remove() int {
	r := mh.data[0]
	mh.data[0] = mh.data[len(mh.data) - 1]
	mh.data = mh.data[:len(mh.data) - 1]
	mh.fix(0)
	return r
}

func (mh *MinHeap) Len() int {
	return len(mh.data)
}


type MaxHeap struct {
	data       []int
}

func NewMaxHeap() *MaxHeap {
	mh := &MaxHeap{
		data: make([]int, 0),
	}
	return mh
}

func (mh *MaxHeap) fix(parent int) bool {
	child := parent * 2 + 1
	updated := false
	for child < len(mh.data) {
		if len(mh.data) > child + 1 && mh.data[child + 1] > mh.data[child] {
			child = child + 1
		}

		if mh.data[parent] > mh.data[child] {
			return updated
		}

		mh.data[parent], mh.data[child] = mh.data[child], mh.data[parent]
		parent = child
		child = parent * 2 + 1
		updated = true
	}
	return updated
}

func (mh *MaxHeap) Init(data []int) {
	mh.data = append(mh.data, data...)
	for i := len(mh.data) / 2 - 1;i >= 0;i -- {
		mh.fix(i)
	}
}

func (mh *MaxHeap) Insert(item int) {
	mh.data = append(mh.data, item)
	pos := len(mh.data) / 2 - 1
	updated := true
	for pos >= 0 && updated == true {
		updated  = mh.fix(pos)
		pos = (pos + 1) / 2 - 1
	}
}

func (mh *MaxHeap) Remove() int {
	r := mh.data[0]
	mh.data[0] = mh.data[len(mh.data) - 1]
	mh.data = mh.data[:len(mh.data) - 1]
	mh.fix(0)
	return r
}

func (mh *MaxHeap) Len() int {
	return len(mh.data)
}
