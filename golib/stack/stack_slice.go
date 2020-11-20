package stack

type Stack_Slice struct {
	items  []int
}

func NewStack_Slice() *Stack_Slice {
	return &Stack_Slice{
		items : make([]int, 0),
	}
}

func (ss *Stack_Slice) Push(value int) {
	ss.items = append(ss.items, value)
}

func (ss *Stack_Slice) Pop() int {
	if ss.Empty() {
		top := ss.items[len(ss.items) - 1]
		ss.items = ss.items[0 : len(ss.items) - 1]
		return top
	}
	return 0
}

func (ss *Stack_Slice) Peak() interface{} {
	if !ss.Empty() {
		return ss.items[len(ss.items) - 1]
	}
	return 0
}

func (ss *Stack_Slice) Len() int {
	return len(ss.items)
}

func (ss *Stack_Slice) Empty() bool {
	return ss.Len() == 0
}
