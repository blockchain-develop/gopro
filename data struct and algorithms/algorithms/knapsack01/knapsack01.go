package knapsack01

import "fmt"

/*
暴力解法
*/


/*
递归解法
*/
func knapsack01_recursive(w []int, v []int, c int, index int) int {
	if index < 0 {
		return 0
	}
	one := knapsack01_recursive(w, v, c, index - 1)
	two := 0
	if w[index] <= c {
		two = knapsack01_recursive(w, v, c-w[index], index - 1) + v[index]
	}
	if one < two {
		return two
	} else {
		return one
	}
}
func Knapsack01_Recursive(w []int, v []int, c int) int {
	return knapsack01_recursive(w, v, c, len(v) - 1)
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
func Knapsack01_DP(w []int, v []int, c int) int {
	x_len, y_len := len(w), c
	keep := make([][]int, y_len + 1)
	for i := 0;i <= y_len;i ++ {
		keep[i] = make([]int, x_len + 1)
	}
	for j := 1;j <= x_len;j ++ {
		for i := 1;i <= y_len;i ++ {
			if keep[i][j-1] > 0 {
				if keep[i][j-1] > keep[i][j] {
					keep[i][j] = keep[i][j-1]
				}
				if i+w[j-1] <= c && keep[i][j-1]+v[j-1] > keep[i+w[j-1]][j] {
					keep[i+w[j-1]][j] = keep[i][j-1] + v[j-1]
				}
			}
		}
		if w[j - 1] <= c && keep[w[j - 1]][j] < v[j - 1]{
			keep[w[j - 1]][j] = v[j - 1]
		}
	}
	// print keep
	y := make([]int, 0)
	for i := 0;i < y_len;i ++ {
		y = append(y, i+1)
	}
	print(w, y, keep)

	max_value := 0
	for i := 0;i < y_len + 1;i ++ {
		if keep[i][x_len] > max_value {
			max_value = keep[i][x_len]
		}
	}
	return max_value
}