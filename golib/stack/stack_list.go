package stack

import "container/list"

type Stack_List struct {
	list *list.List
}

func NewStack_List() *Stack_List {
	return &Stack_List{list:list.New()}
}

func (sl *Stack_List) Push(value interface{}) {
	sl.list.PushBack(value)
}

func (sl *Stack_List) Pop() interface{} {
	e := sl.list.Back()
	if e != nil {
		sl.list.Remove(e)
		return e.Value
	}
	return nil
}

func (sl *Stack_List) Peak() interface{} {
	e := sl.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (sl *Stack_List) Len() int {
	return sl.list.Len()
}

func (sl *Stack_List) Empty() bool {
	return sl.Len() == 0
}
