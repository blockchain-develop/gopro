package knapsackn

import "fmt"

/*
暴力解法
*/


/*
递归解法
*/
func knapsackn_recursive(w []int, v []int, n []int, c int, index int) int {
	if index < 0 {
		return 0
	}
	one := knapsackn_recursive(w, v, n, c, index - 1)
	two := 0
	for i := 1;i <= n[index];i ++ {
		if w[index] * i <= c {
			two = knapsackn_recursive(w, v, n, c - w[index] * i, index-1) + v[index] * i
		} else {
			break
		}
	}
	if one < two {
		return two
	} else {
		return one
	}
}
func KnapsackN_Recursive(w []int, v []int, n []int, c int) int {
	return knapsackn_recursive(w, v, n, c, len(v) - 1)
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
func KnapsackN_DP(w []int, v []int, n []int, c int) int {
	x_len, y_len := len(w), c
	keep := make([][]int, y_len + 1)
	for i := 0;i <= y_len;i ++ {
		keep[i] = make([]int, x_len + 1)
	}
	for j := 1;j <= x_len;j ++ {
		for i := 1;i <= y_len;i ++ {
			if keep[i][j-1] > 0 {
				for m := 0;m <= n[j-1];m ++ {
					if i + w[j-1]*m <= c && keep[i + w[j-1]*m][j] < keep[i][j-1] + v[j-1]*m {
						keep[i+w[j-1]*m][j] = keep[i][j-1] + v[j-1]*m
					}
				}
			}
		}
		for i := 1;i <= n[j - 1];i ++ {
			if w[j-1] * i <= c && keep[w[j-1]*i][j] < v[j-1]*i {
				keep[w[j-1]*i][j] = v[j-1]*i
			}
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