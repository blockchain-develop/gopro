package shortestpath

import (
	"fmt"
	"testing"
)

func TestShortestPath_Dijkstra(t *testing.T) {
	graph := [][]int{
		[]int{ 0, 4, 0, 0, 0, 0, 0, 8, 0 },
		[]int{ 4, 0, 8, 0, 0, 0, 0, 11, 0 },
		[]int{ 0, 8, 0, 7, 0, 4, 0, 0, 2 },
		[]int{ 0, 0, 7, 0, 9, 14, 0, 0, 0 },
		[]int{ 0, 0, 0, 9, 0, 10, 0, 0, 0 },
		[]int{ 0, 0, 4, 14, 10, 0, 2, 0, 0 },
		[]int{ 0, 0, 0, 0, 0, 2, 0, 1, 6 },
		[]int{ 8, 11, 0, 0, 0, 0, 1, 0, 7 },
		[]int{ 0, 0, 2, 0, 0, 0, 6, 7, 0 } }
	dist, path := ShortestPath_Dijkstra(graph)
	for i := 0;i < 9;i ++ {
		fmt.Printf("dist and path from 0 to %d, dist: %d, path: [ ", i, dist[i])
		for j := 0;j < len(path[i]);j ++ {
			fmt.Printf("%d ", path[i][j])
		}
		fmt.Print("]\n")
	}
}
