package dicethrow

/*
朴素解法：
通过M进制数的方式遍历所有可能的组合，判断组合是否满足要求
 */
func pow(x int, y int) int {
	z := 1
	for i := 0;i < y;i ++ {
		z *= x
	}
	return z
}

func unit(x int, y int, u int) []int {
	r := make([]int, u)
	for i := 0;i < u;i ++ {
		r[i] = x % y
		x = x / y
	}
	return r
}

func DiceThrow_Naive(n int, m int, x int) int {
	counter := 0
	all := pow(m, n)
	for i := 0;i < all;i ++ {
		u := unit(i, m, n)
		sum := 0
		for _,j := range u {
			sum += (j+1)
		}
		if sum == x {
			counter ++
		}
	}
	return counter
}

/*
递归解法：
find X from n dice is : Sum(m, n, X)
Sum(m, n, X) =
	Sum(m, n - 1, X - 1) （take 1 from n dice) +
	Sum(m, n - 1, X - 2) （take 2 from n dice) +
	Sum(m, n - 1, X - 3) （take 3 from n dice) +
    ......
*/
func dicethrow_recursive(m int, n int, x int) int {
	if x < 0 {
		return 0
	}
	if n == 0 && x != 0 {
		return 0
	}
	if n ==0 && x == 0 {
		return 1
	}
	counter := 0
	for i := 1;i <= m;i ++ {
		counter += dicethrow_recursive(m, n - 1, x - i)
	}
	return counter
}
func DiceThrow_Recursive(m int, n int, x int) int {
	return dicethrow_recursive(m, n, x)
}

/*
动态规划解法：
*/
func DiceThrow_DP(m int, n int, x int) int {
	keep := make([][]int, x + 1)
	for i := 0;i <= x;i ++ {
		keep[i] = make([]int, n + 1)
	}
	for i := 1;i <= m && i <= x;i ++ {
		keep[i][1] = 1
	}
	for j := 2;j <= n;j ++ {
		for i := 1; i <= x; i ++ {
			for k := 1;k <= m && i + k <= x;k ++ {
				keep[i + k][j] += keep[i][j-1]
			}
		}
	}
	return keep[x][n]
}
