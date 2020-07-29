package partitionset
/*
朴素解法
*/
func partitionset(src []int, check int) ([]int, []int) {
	checka := make([]int, 0)
	checkb := make([]int, 0)
	index := 0
	very := 1
	for index < len(src) {
		if very & check > 0 {
			checka = append(checka, src[index])
		} else {
			checkb = append(checkb, src[index])
		}
		very = very << 1
		index ++
	}
	return checka, checkb
}
func setdiff(a []int, b[]int) int {
	totala, totalb := 0, 0
	for _, item := range a {
		totala += item
	}
	for _, item := range b {
		totalb += item
	}
	if totala > totalb {
		return totala - totalb
	} else {
		return totalb - totala
	}
}
func PartitionSet_Naive(src []int) ([]int, []int) {
	src_len := len(src)
	power := 1 << uint32(src_len)
	diff := 100000
	resulta, resultb := make([]int, 0), make([]int, 0)
	for i := 0;i < power;i ++ {
		seta, setb := partitionset(src, i)
		diff_temp := setdiff(seta, setb)
		if diff_temp < diff {
			diff = diff_temp
			resulta = seta
			resultb = setb
		}
	}
	return resulta, resultb
}

/*
递归解法
*/
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func abs(a int) int {
	if a < 0 {
		return 0-a
	} else {
		return a
	}
}
func findMinDiff(src []int, i int, sumCalculated int, sumTotal int) int {
	if i == 0 {
		return abs(sumTotal-sumCalculated-sumCalculated)
	}
	return min(findMinDiff(src, i-1, sumCalculated+src[i-1], sumTotal),
	findMinDiff(src, i-1, sumCalculated, sumTotal))
}
func PartitionSet_Recursive(src []int) int {
	sumTotal := 0
	for _, item := range src {
		sumTotal += item
	}
	return findMinDiff(src, len(src), 0, sumTotal)
}
/*
动态规划解法
*/
func PartitionSet_DP(src []int) (int) {
	src_len := len(src)
	src_sum := 0
	for _, item := range src {
		src_sum += item
	}
	keep := make([][]byte, src_len+1)
	for i := 0;i <= src_len;i ++ {
		keep[i] = make([]byte, src_sum+1)
	}
	for j := 0;j <= src_sum;j ++ {
		keep[0][j] = 0
	}
	for i := 0;i <= src_len;i ++ {
		keep[i][0] = 1
	}
	for i := 1;i <= src_len;i ++ {
		for j := src_sum;j >= 0;j -- {
			keep[i][j] = keep[i-1][j]
			if j + src[i-1] <= src_sum {
				keep[i][j+src[i-1]] |= keep[i-1][j]
			}
		}
	}
	index := 0
	sumCalculated := 0
	for i := 0;i < src_sum/2;i ++ {
		if keep[src_len][src_sum/2+index] != 0 {
			sumCalculated = src_sum/2+index
			break
		} else if keep[src_len][src_sum/2-index] != 0 {
			sumCalculated = src_sum/2-index
			break
		} else {
			index ++
		}
	}
	return abs(src_sum - sumCalculated - sumCalculated)
}