package middlenum4sortedarrays

/*
func middlenum4sortedarrays(a []int, b []int, left int, right int) float32 {
	size_a, size_b := len(a), len(b)
	index_a1, index_a2 := (size_a - 1)/2, size_a/2
	index_b1, index_b2 := (size_b - 1)/2 , size_b/2
	middle_a := (float32(a[index_a1]) + float32(a[index_a2]))/2
	middle_b := (float32(a[index_b1]) + float32(a[index_b2]))/2

	if middle_a >= middle_b {
		return middlenum4sortedarrays(a[0:index_a1+1], b[index_b2: len(b)], diff + len(a) - index_a2 - index_b1)
	} else {
		return middlenum4sortedarrays(a[index_a2:len(a)], b[0: index_b1+1], diff - index_a1 + len(b) - index_b2)
	}
}
func MiddleNum4SortedArrays(a []int, b []int) float32 {
	return middlenum4sortedarrays(a, b, 0)
}
*/