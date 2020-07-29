package fibonacci

/*
朴素算法
*/
func Fibonacci_Naive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci_Naive(n - 1) + Fibonacci_Naive(n - 2)
}

/*
迭代法
*/
func Fibonacci_Tterative(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	prepre := 0
	pre := 1
	cur := 0
	for i := 2;i <= n;i ++ {
		cur = prepre + pre
		prepre = pre
		pre = cur
	}
	return cur
}