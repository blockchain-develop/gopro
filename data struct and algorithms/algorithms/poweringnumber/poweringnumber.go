package poweringnumber

/*
朴素算法
*/
func Powering_Naive(m int, n int) int {
	power := 1
	for i := 0;i < n;i ++ {
		power *= m
	}
	return power
}

/*
分治算法
*/
func Powering_DivideConquer(m int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return m
	}
	return Powering_DivideConquer(m, n/2) * Powering_DivideConquer(m, (n+1)/2)
}