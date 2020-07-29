package subsequencesearch

func SubsequenceSearch_Force(data []byte, target []byte) int {
	sourceLen := len(data)
	targetLen := len(target)
	if sourceLen < targetLen {
		return -1
	}
	for i:= 0;i <= sourceLen - targetLen;i ++ {
		j := 0
		for j = 0;j < targetLen;j ++ {
			if data[i + j] != target[j] {
				break
			}
		}
		if j >= targetLen {
			return i
		}
	}
	return -1
}

func pie(p []byte) []int {
	a := make([]int, len(p))
	a[0] = 0
	k := 0
	for q := 1;q < len(p);q ++ {

	}
}

func SubsequenceSearch_KMP(data []byte, target []byte) int {

}
