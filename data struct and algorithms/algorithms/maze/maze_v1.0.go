package maze

import "fmt"

/*
1. 栈
2. 寻找一条路劲
*/

type stack struct {
	items []*pos
}

func NewStack() *stack {
	s := &stack{
		items : make([]*pos, 0),
	}
	return s
}

func (s *stack) Push(item *pos) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() *pos {
	item := s.items[len(s.items) - 1]
	s.items = s.items[0 : len(s.items) - 1]
	return item
}

func (s *stack) Peak() *pos {
	item := s.items[len(s.items) - 1]
	return item
}

func (s *stack) Empty() bool {
	return len(s.items) == 0
}

func (s *stack) Print() {
	for _, item := range s.items {
		fmt.Printf("(%d, %d), ", item.x, item.y)
	}
	fmt.Printf("\n")
}


type pos struct {
	x int
	y int
}

func nextStep(maze [][]int, current *pos) *pos {
	if maze[current.x + 1][current.y] == 0 {
		return &pos{
			x: current.x + 1,
			y: current.y,
		}
	} else if maze[current.x][current.y + 1] == 0 {
		return &pos{
			x: current.x,
			y: current.y + 1,
		}
	} else if maze[current.x - 1][current.y] == 0 {
		return &pos{
			x: current.x - 1,
			y: current.y,
		}
	} else if maze[current.x][current.y- 1] == 0 {
		return &pos{
			x: current.x,
			y: current.y - 1,
		}
	}
	return nil
}

func FindPath_v10(maze [][]int, in_x int, in_y int, out_x int, out_y int) {
	stack := NewStack()
	current := &pos {
		x: in_x,
		y: in_y,
	}
	out := &pos {
		x: out_x,
		y: out_y,
	}
	for current.x != out.x || current.y != out.y {
		next := nextStep(maze, current)
		if next != nil {
			stack.Push(current)
			current = next
			maze[current.x][current.y] = 1
		} else {
			if stack.Empty() {
				return
			}
			current = stack.Pop()
		}
	}
	stack.Print()
}