package mergesort

func merge(data []int32, start int, end int, middle int) {
	left := start
	right := middle + 1
	which := make([]int32, end - start + 1)
	for left <= middle || right <= end {
		if left <= middle && right <= end {
			if data[left] <= data[right] {
				which[left-start+right-middle-1] = data[left]
				left ++
			} else {
				which[left-start+right-middle-1] = data[right]
				right ++
			}
		} else if left <= middle {
			which[left-start+right-middle-1] = data[left]
			left ++
		} else {
			which[left-start+right-middle-1] = data[right]
			right ++
		}
	}

	for i := 0;i < end - start + 1;i ++ {
		data[start + i] = which[i]
	}
}
func mergesort(data []int32, start int, end int, k int) {
	if end - start == 0 {
		return
	}

	index := start
	round := k
	for round > 0 {
		slice := ((end - index + 1) + (round-1))/round
		mergesort(data, index, index + slice - 1, k)
		round --
		index = index + slice
	}
	merge(data, start, end, (end + start) / 2)
}

func MergeSort(data []int32) {
	mergesort(data, 0, len(data) - 1, 2)
}

