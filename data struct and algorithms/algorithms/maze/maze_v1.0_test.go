package maze

import "testing"

func TestMazev10(t *testing.T) {
	max_x := 8
	max_y := 8
	maze := make([][]int, max_x + 2)
	for i, _ := range maze {
		maze[i] = make([]int, max_y + 2)
	}
	for i := 0;i < max_x + 2;i ++ {
		maze[i][0] = 1
		maze[i][max_y + 1] = 1
	}
	for j := 0;j < max_y + 2;j ++ {
		maze[0][j] = 1
		maze[max_x + 1][j] = 1
	}
	FindPath_v10(maze, 1, 1, max_x, max_y)
}
