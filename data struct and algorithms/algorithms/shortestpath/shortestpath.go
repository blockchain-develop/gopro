package shortestpath

import (
	"github.com/gopro/datastruct/linear/queue"
	"math"
)

func ShortestPath_Dijkstra(graph [][]int) ([]int, [][]int) {
	v := len(graph)
	dist := make([]int, v)
	path := make([][]int, v)
	for i := 0;i < v;i ++ {
		path[i] = make([]int, 0)
	}
	for i := 0;i < v;i ++ {
		dist[i] = math.MaxInt32
	}
	dist[0] = 0
	path[0] = append(path[0], 0)
	q := queue.NewQueue()
	q.Add(0)

	for true {
		item := q.Remove()
		if item == -1 {
			break
		}
		for j := 0;j < v;j ++ {
			if graph[item][j] != 0 && dist[item] + graph[item][j] < dist[j] {
				dist[j] = dist[item] + graph[item][j]
				path[j] = append(path[j][:0], path[item]...)
				path[j] = append(path[j], j)
				q.Add(j)
			}
		}
	}
	return dist, path
}
