package minimumspanningtrees

import (
	"fmt"
	"testing"
)

func TestMinimumSpanningTrees_Prim(t *testing.T) {
	graph := [][]int{[]int{ 0, 4, 0, 0, 0, 0, 0, 8, 0 },
		[]int{ 4, 0, 8, 0, 0, 0, 0, 11, 0 },
		[]int{ 0, 8, 0, 7, 0, 4, 0, 0, 2 },
		[]int{ 0, 0, 7, 0, 9, 14, 0, 0, 0 },
		[]int{ 0, 0, 0, 9, 0, 10, 0, 0, 0 },
		[]int{ 0, 0, 4, 14, 10, 0, 2, 0, 0 },
		[]int{ 0, 0, 0, 0, 0, 2, 0, 1, 6 },
		[]int{ 8, 11, 0, 0, 0, 0, 1, 0, 7 },
		[]int{ 0, 0, 2, 0, 0, 0, 6, 7, 0 } };
	tree := MinimumSpanningTrees_Prim(graph)
	degree_search := make([]int, 0)
	tree.DegreeFirstSearch(&degree_search)
	fmt.Printf("degree search, [ ")
	for j := 0;j < len(degree_search);j ++ {
		fmt.Printf("%d ", degree_search[j])
	}
	fmt.Print("]\n")

	depth_search := make([]int, 0)
	tree.DepthFirstSearch(&depth_search)
	fmt.Printf("depth search, [ ")
	for j := 0;j < len(depth_search);j ++ {
		fmt.Printf("%d ", depth_search[j])
	}
	fmt.Print("]\n")
}

