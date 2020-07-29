package subsetsum

import "fmt"

/*
暴力解法
*/


/*
递归解法
*/
func subsetsum_recrusive(x []int, sum int, index int) bool {
	if sum == 0 {
		return true
	}
	if sum < 0 || index < 0 {
		return false
	}
	return subsetsum_recrusive(x, sum-x[index], index-1) || subsetsum_recrusive(x, sum, index-1)
}
func SubsetSum_Recursive(x []int, sum int) bool {
	return subsetsum_recrusive(x, sum, len(x) - 1)
}

/*
动态规划解法
*/
func print(x []int, y []int, keep [][]int) {
	fmt.Printf("\t\t")
	for i := 1;i <= len(x);i ++ {
		fmt.Printf("\t%d", (x[i-1]))
	}
	fmt.Printf("\n")
	for i, item := range keep {
		if i > 0 {
			fmt.Printf("\t%d", (y[i-1]))
		} else {
			fmt.Printf("\t")
		}
		for _, value := range item {
			fmt.Printf("\t%d", value)
		}
		fmt.Printf("\n")
	}
}
func SubsetSum_DP(x []int, sum int) bool {
	x_len, y_len := len(x), sum
	keep := make([][]int, y_len + 1)
	for i := 0;i <= y_len;i ++ {
		keep[i] = make([]int, x_len + 1)
	}
	for j := 1;j <= x_len;j ++ {
		for i := 1;i <= y_len;i ++ {
			if keep[i][j-1] == 1 {
				keep[i][j] = 1
				if i + x[j-1] <= y_len {
					keep[i+x[j-1]][j] = 1
				}
			}
		}
		if x[j-1] <= y_len {
			keep[x[j-1]][j] = 1
		}
	}
	//
	y := make([]int, 0)
	for i := 0;i < y_len;i ++ {
		y = append(y, i+1)
	}
	print(x, y, keep)
	return keep[y_len][x_len] == 1
}