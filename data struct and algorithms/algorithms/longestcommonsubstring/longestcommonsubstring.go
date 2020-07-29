package longestcommonsubstring

func min(a int, b int) int{
	if a <= b {
		return a
	} else {
		return b
	}
}
/*
朴素解法
*/
func LongestCommonSubstring_Naive(s []byte, t []byte) (si, ti, length int) {
	sl, tl := len(s), len(t)
	si, ti, length = -1, -1, 0
	for i := 0;i < sl;i ++ {
		for j := 0;j < tl;j ++ {
			k := 0
			for k = 0;k < min(sl-i, tl-j);k ++ {
				if s[k + i] != t[k + j] {
					break
				}
			}
			if k > length {
				si, ti, length = i, j, k
			}
		}
	}
	return si, ti, length
}

/*
动态规划解法
*/
func LongestCommonSubstring_DP(s []byte, t []byte) (si, ti, length int) {
	sl, tl := len(s), len(t)
	keep := make([][]int, tl + 1)
	for i := range keep {
		keep[i] = make([]int, sl + 1)
	}
	for i := 0;i < sl;i ++ {
		for j := 0;j < tl;j ++ {
			if s[i] == t[j] {
				keep[j+1][i+1] = keep[j][i] + 1
			}
		}
	}
	si,ti = -1, -1
	length = 0
	for i := 1; i < sl;i ++ {
		for j := 1;j < tl;j ++ {
			if keep[j][i] > length {
				si,ti = i,j
				length = keep[j][i]
			}

		}
	}
	return si, ti, length
}