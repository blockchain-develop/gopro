package longestpalindromicsubstring

/*
first, by force
找出所有的子串，判断其是否为回文子串
*/
func LongestPalindromicSubstring_Force(data []byte) (sub []byte) {
	sub = data[0:1]
	for i := 0;i < len(data); i ++ {
		for j := len(data) - 1;j >= i + len(sub);j -- {
			m,n := 0,0
			for m,n = i,j;m < n;m,n = m+1,n-1 {
				if data[m] != data[n] {
					break
				}
			}
			if m >= n {
				sub = data[i:j+1]
			}
		}
	}
	return sub
}
/*
second naive
*/

func LongestPalindromicSubstring_Naive(data []byte) (index, length int) {
	index, length = -1, 0
	for i := 0;i < len(data);i ++ {
		//
		m,n := 0,0
		for m,n = i,i+1;m >= 0 && n < len(data);m,n = m - 1, n + 1 {
			if data[m] != data[n] {
				break
			}
		}
		if n - m - 1 > length {
			length = n - m -1
			index = m + 1
		}
		//
		for m,n = i-1,i+1;m >= 0 && n < len(data);m,n = m - 1, n + 1 {
			if data[m] != data[n] {
				break
			}
		}
		if n - m - 1 > length {
			length = n - m -1
			index = m + 1
		}
	}
	return index, length
}

/*
动态规划解法
*/
func LongestPalindromicSubstring_DP(data []byte) int {
	dataLen := len(data)
	keep := make([][]bool, dataLen)
	for i := 0;i < dataLen;i ++ {
		keep[i] = make([]bool, dataLen)
	}
	for i := 0;i < dataLen;i ++ {
		keep[i][i] = true
	}
	for i := dataLen - 2;i >= 0;i -- {
		for j := i + 1;j < dataLen;j ++ {
			if data[i] == data[j] {
				if i + 1 >= j - 1 {
					keep[i][j] = true
				} else if keep[i + 1][j - 1] == true {
					keep[i][j] = true
				} else {
					keep[i][j] = false
				}
			} else {
				keep[i][j] = false
			}
		}
	}
	resultLen := 0
	for i := 0;i < dataLen;i ++ {
		for j := 0;j < dataLen;j ++ {
			if j + i < dataLen && keep[j][j + i] == true {
				resultLen = i + 1
			}
		}
	}
	return resultLen
}

