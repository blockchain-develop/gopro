package towerofhanoi

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type HanoiStack struct {
	stack *stack.Stack
	num byte
}
func NewHanoiStack(num byte) *HanoiStack {
	stack := &HanoiStack{
		stack : stack.New(),
		num : num,
	}
	return stack
}
func (stack *HanoiStack)Push(item int) {
	stack.stack.Push(item)
}
func (stack *HanoiStack)Pop() int {
	item := stack.stack.Pop()
	return item.(int)
}

func towerofhanoi(a,b,c *HanoiStack, size int) {
	if size > 1 {
		towerofhanoi(a, c, b, size-1)
	}
	which := a.Pop()
	c.Push(which)
	fmt.Printf("%c -> %c, %d\n", a.num, c.num, which)
	if size > 1 {
		towerofhanoi(b, a, c, size-1)
	}
}

func TowerofHanoi(height int) {
	stacka := NewHanoiStack('a')
	stackb := NewHanoiStack('b')
	stackc := NewHanoiStack('c')
	for i := height;i >=1;i -- {
		stacka.Push(i)
	}
	towerofhanoi(stacka, stackb, stackc, height)
}
