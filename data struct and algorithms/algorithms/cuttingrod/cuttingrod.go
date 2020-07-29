package cuttingrod

import "fmt"

/*
递归解法
*/
func cuttingrod_recursive(l []int, v []int, index int, rodl int) int {
	max_value := 0
	i := 0
	for true {
		if index < 0 {
			break
		}
		if rodl - l[index] * i < 0 {
			break
		}
		result := cuttingrod_recursive(l, v, index - 1, rodl - l[index] * i) + v[index] * i
		if result > max_value {
			max_value = result
		}
		i++
	}
	return max_value
}
func CuttingRod_Recursive(l []int, v []int, rodl int) int {
	return cuttingrod_recursive(l, v, len(v) - 1, rodl)
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
func CuttingRod_DP(l []int, v []int, rodl int) int {
	x_len, y_len := len(l), rodl
	keep := make([][]int, y_len + 1)
	for i := 0;i <= y_len;i ++ {
		keep[i] = make([]int, x_len + 1)
	}
	for j := 1;j <= x_len;j ++ {
		for i := 1;i <= y_len;i ++ {
			if keep[i][j-1] > 0 {
				m := 0
				for true {
					if i + l[j-1]*m <= rodl {
						if keep[i][j-1] + v[j-1]*m > keep[i+l[j-1]*m][j] {
							keep[i+l[j-1]*m][j] = keep[i][j-1] + v[j-1]*m
						}
						m ++
					} else {
						break
					}
				}
			}
		}
		m := 1
		for true {
			if l[j-1]*m <= rodl {
				if keep[l[j-1]*m][j] < v[j-1]*m {
					keep[l[j-1]*m][j] = v[j-1]*m
				}
				m ++
			} else {
				break
			}
		}
	}
	// print keep
	y := make([]int, 0)
	for i := 0;i < y_len;i ++ {
		y = append(y, i+1)
	}
	print(l, y, keep)
	//
	max_value := 0
	for i := 0;i < y_len + 1;i ++ {
		if keep[i][x_len] > max_value {
			max_value = keep[i][x_len]
		}
	}
	return max_value
}