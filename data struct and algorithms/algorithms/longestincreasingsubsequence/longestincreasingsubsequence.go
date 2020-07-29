package longestincreasingsubsequence

import (
	"fmt"
	"sort"
)

/*
朴素解法1， 求所有的子序列，找出最长的递增子序列
*/
func getSubsequence(src []int, check int) []int {
	checked := make([]int, 0)
	index := 0
	very := 1
	for very <= check && index < len(src) {
		if very & check > 0 {
			checked = append(checked, src[index])
		}
		very = very << 1
		index ++
	}
	return checked
}
func isIncreasing(src []int) bool {
	prev := -100000
	for _, item := range src {
		if item < prev {
			return false
		}
		prev = item
	}
	return true
}
func LongestIncreasingSubsequence_Naive(src []int) []int {
	src_len := len(src)
	power := 1 << uint32(src_len)
	longest := 0
	longestincreasingsubsequence := make([]int, 0)
	for i := 0;i < power;i ++ {
		sub := getSubsequence(src, i)
		valid := isIncreasing(sub)
		if valid == true && len(sub) > longest {
			longest = len(sub)
			longestincreasingsubsequence = sub
		}
	}
	return longestincreasingsubsequence
}

/*
递归解法
*/
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func check(src []int) bool {
	prev := 10000000
	for _,item := range src {
		if item > prev {
			return false
		}
		prev = item
	}
	return true
}
func getSubsequence_Recursive(src []int, i int, oldkeep []int) []int {
	if i == 0 {
		valid := check(oldkeep)
		if valid == false {
			return nil
		} else {
			return oldkeep
		}
	}
	newkeep := make([]int, len(oldkeep))
	copy(newkeep, oldkeep)
	newkeep = append(newkeep, src[i-1])
	sub1 := getSubsequence_Recursive(src, i-1, newkeep)
	sub2 := getSubsequence_Recursive(src, i-1, oldkeep)
	if sub1 == nil && sub2 == nil {
		return nil
	} else if sub1 == nil {
		return sub2
	} else if sub2 == nil {
		return sub1
	} else if len(sub1) >= len(sub2) {
		return sub1
	} else {
		return sub2
	}
}
func LongestIncreasingSubsequence_Recursive(src []int) []int {
	newkeep := make([]int, 0)
	return getSubsequence_Recursive(src, len(src), newkeep)
}

/*
递归解法2
*/
func longestIncreasingSubsequence_Recursive2(src []int, i int) int {
	max_ending_here := 1
	res := 1
	for j := 0;j < i;j ++ {
		if src[j] < src[i] {
			res = longestIncreasingSubsequence_Recursive2(src, j)
			if res + 1 > max_ending_here {
				max_ending_here = res + 1
			}
		}
	}
	return max_ending_here
}

func LongestIncreasingSubsequence_Recursive2(src []int) int {
	return longestIncreasingSubsequence_Recursive2(src, len(src) - 1)
}

/*
动态规划解法
*/
func LongestIncreasingSubsequence_DP(src []int) int {
	keep := make([]int, len(src))
	keep[0] = 1
	for i := 1;i < len(src);i ++ {
		longest := 1
		for j := 0;j < i;j ++ {
			if src[i] > src[j] {
				if keep[j] + 1 > longest {
					longest = keep[j] + 1
				}
			}
		}
		keep[i] = longest
	}
	fmt.Printf("%v\n", keep)
	longest := 0
	for i := 0;i < len(keep);i ++ {
		if keep[i] > longest {
			longest = keep[i]
		}
	}
	return longest
}

/*
最长公共子串解法
*/
func LongestCommonSubsequence(x []int, y[]int) int {
	x_len, y_len := len(x), len(y)
	keep := make([][]int, y_len + 1)
	for i := 0;i <= y_len;i ++ {
		keep[i] = make([]int, x_len + 1)
	}
	for i := 1;i <= y_len;i ++ {
		for j := 1;j <= x_len;j ++ {
			if x[j - 1] == y[i - 1] {
				keep[i][j] = keep[i-1][j-1] + 1
			} else {
				keep[i][j] = max(keep[i-1][j], keep[i][j-1])
			}
		}
	}
	return keep[y_len][x_len]
}
func LongestIncreasingSubsequence_New(src []int) int {
	sortedSrc := make([]int, len(src))
	copy(sortedSrc, src)
	sort.Slice(sortedSrc, func(i, j int) bool {
		return sortedSrc[i] < sortedSrc[j]
	})
	return LongestCommonSubsequence(sortedSrc, src)
}