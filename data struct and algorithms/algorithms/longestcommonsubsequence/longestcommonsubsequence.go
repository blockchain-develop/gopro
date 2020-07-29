package longestcommonsubsequence

import "fmt"

/*
暴力解法
*/
func getSubsequenceByBit(x []byte, b int) []byte {
	z := make([]byte, 0)
	index := 0
	for b > 0 && index < len(x) {
		if b & 1 > 0 {
			z = append(z, x[index])
		}
		b = b >> 1
		index ++
	}
	return z
}
func isSubsequence(z []byte, y[]byte) bool {
	i, j := 0, 0
	for i,j = 0,0;i < len(z) && j < len(y); {
		if z[i] == y[j] {
			i ++
			j ++
		} else {
			j ++
		}
	}
	if i == len(z) {
		return true
	} else {
		return false
	}
}
func LongestCommonSubsequence_Naive(x []byte, y []byte) (int, [][]byte) {
	x_len := len(x)
	x_power := 1 << uint(x_len)
	longest := 0
	longestSubsequence := make([][]byte, 0)
	for i := 1;i < x_power;i ++ {
		z := getSubsequenceByBit(x, i)
		isSub := isSubsequence(z, y)
		if isSub == false {
			continue
		}
		if len(z) > longest {
			longestSubsequence = make([][]byte, 0)
			longestSubsequence = append(longestSubsequence, z)
			longest = len(z)
		} else if len(z) == longest {
			longestSubsequence = append(longestSubsequence, z)
		} else {

		}
	}
	return longest, longestSubsequence
}

/*
递归解法
*/
var (
	longestsub = make([][]byte, 0)
	longest = 0
)
func addsubsequence(newsubsequence []byte){
	newlen := len(newsubsequence)
	if newlen < longest {
		return
	}
	if newlen == longest {
		longestsub = append(longestsub, newsubsequence)
		return
	}
	longestsub = longestsub[0:0]
	longestsub = append(longestsub, newsubsequence)
	longest = len(newsubsequence)
}
func longestcommonsubsequence_recursive(x []byte, i int, oldkeep []byte, y[]byte) {
	if i == len(x) {
		valid := isSubsequence(oldkeep, y)
		if valid == false {
			return
		} else {
			addsubsequence(oldkeep)
			return
		}
	}
	newkeep := make([]byte, len(oldkeep))
	copy(newkeep, oldkeep)
	newkeep = append(newkeep, x[i])
	longestcommonsubsequence_recursive(x, i+1, newkeep, y)
	longestcommonsubsequence_recursive(x, i+1, oldkeep, y)
}
func LongestCommonSubsequence_Recursive(x []byte, y []byte) (int, [][]byte) {
	newkeep := make([]byte, 0)
	longestcommonsubsequence_recursive(x, 0, newkeep, y)
	return longest, longestsub
}

/*
动态规划解法
*/
func max(x int, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
func print(x []byte, y []byte, keep [][]int) {
	fmt.Printf("%s","    ")
	for i := 1;i <= len(x);i ++ {
		fmt.Printf("%s ", string(x[i-1]))
	}
	fmt.Printf("\n")
	for i, item := range keep {
		if i > 0 {
			fmt.Printf("%s ", string(y[i-1]))
		} else {
			fmt.Printf("  ")
		}
		for _, value := range item {
			fmt.Printf("%d ", value)
		}
		fmt.Printf("\n")
	}
}
func LongestCommonSubsequence_DP(x []byte, y[]byte) int {
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
	// print keep
	print(x, y, keep)
	return keep[y_len][x_len]
}