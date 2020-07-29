package nqueens

/*
暴力解法
*/
func intarraynext(a []int32, b int32) bool {
	i := 0
	for true {
		if i >= len(a) {
			return false
		}
		a[i] = a[i] << 1
		if a[i] >= b {
			a[i] = 1
			i ++
		} else {
			return true
		}
	}
	return false
}
func intarraycheck(a []int32, b int32) bool {
	checkvalue := 1
	for checkvalue < int(b) {
		checktotal := 0
		for _, item := range a {
			if int(item) & checkvalue > 0 {
				checktotal ++
			}
		}
		if checktotal > 1 {
			return false
		}
		checkvalue = checkvalue << 1
	}
	for i := 0;i < len(a) - 1;i ++ {
		checka, checkb := a[i], a[i]
		for j := i + 1;j < len(a);j ++ {
			checka = checka << 1
			checkb = checkb >> 1
			if checka & a[j] > 0 || checkb & a[j] > 0 {
				return false
			}
		}
	}
	return true
}
func NQueens(number int) [][]int32 {
	result := make([][]int32, 0)
	keep := make([]int32, 0)
	for i := 0;i < number;i ++ {
		keep = append(keep, 1)
	}
	power := int32(1 << uint32(number))
	next := true
	for next == true {
		valid := intarraycheck(keep, power)
		if valid {
			newkeep := make([]int32, len(keep))
			copy(newkeep, keep)
			result = append(result, newkeep)
		}
		next = intarraynext(keep, power)
	}
	return result
}
