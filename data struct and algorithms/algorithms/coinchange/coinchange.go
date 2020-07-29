package coinchange

import "fmt"

/*
暴力解法
*/


/*
递归解法
*/
func coinchange_recursive(coins []int, index int, amount int) int {
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}
	if index < 0 {
		return 0
	}
	number := 0
	m := 0
	for true {
		if coins[index]*m <= amount {
			result := coinchange_recursive(coins, index-1, amount-coins[index]*m)
			m ++
			number += result
		} else {
			break
		}
	}
	return number
}
func CoinChange_Recursive(coins []int, amount int) int {
	return coinchange_recursive(coins, len(coins)-1, amount)
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
func CoinChange_DP(coins []int, amount int) int {
	x_len, y_len := len(coins), amount
	keep := make([][]int, y_len + 1)
	for i := 0;i <= y_len;i ++ {
		keep[i] = make([]int, x_len + 1)
	}
	for j := 1;j <= x_len;j ++ {
		for i := 1;i <= y_len;i ++ {
			if keep[i][j-1] > 0 {
				m := 0
				for true {
					if i + coins[j-1]*m <= amount {
						keep[i + coins[j-1]*m][j] += keep[i][j-1]
						m ++
					} else {
						break
					}
				}
			}
		}
		m := 1
		for true{
			if coins[j-1]*m <= amount {
				keep[coins[j-1]*m][j] += 1
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
	print(coins, y, keep)
	return keep[y_len][x_len]
}