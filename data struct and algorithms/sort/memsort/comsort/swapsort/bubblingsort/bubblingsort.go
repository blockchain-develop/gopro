package bubblingsort

func BubblingSort(data []int32) {
	for i := len(data) - 1;i > 0;i -- {
		for j := 0;j < i;j ++ {
			if data[j] > data[j + 1] {
				key := data[j]
				data[j] = data[j + 1]
				data[j + 1] = key
			}
		}
	}
}
