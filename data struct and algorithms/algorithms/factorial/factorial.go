package factorial

/*
Iterative
*/

func Fac_Tterative(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	fac := 1
	for i := 2;i <= n;i ++ {
		fac *= i
	}
	return fac
}

/*
Recursive
*/
func Fac_Recursive(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * Fac_Recursive(n - 1)
}
