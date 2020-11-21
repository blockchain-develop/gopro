package maze

import "fmt"

/*
1. 队列
2. 寻找一条最短路劲
*/

type queue struct {
	items []*pos
}

func NewQueue() *queue {
	s := &queue{
		items : make([]*pos, 0),
	}
	return s
}

func (s *queue) Push(item *pos) {
	s.items = append(s.items, item)
}

func (s *queue) Pop() *pos {
	item := s.items[0]
	s.items = s.items[1 : len(s.items) - 1]
	return item
}

func (s *queue) Peak() *pos {
	item := s.items[0]
	return item
}

func (s *queue) Empty() bool {
	return len(s.items) == 0
}

func (s *queue) Print() {
	for _, item := range s.items {
		fmt.Printf("(%d, %d), ", item.x, item.y)
	}
	fmt.Printf("\n")
}

func FindPath_v20(maze [][]int, in_x int, in_y int, out_x int, out_y int) {
	queue := NewQueue()
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
			queue.Push(current)
			current = next
			maze[current.x][current.y] = 1
		} else {
			if queue.Empty() {
				return
			}
			current = queue.Pop()
		}
	}
	queue.Print()
}
