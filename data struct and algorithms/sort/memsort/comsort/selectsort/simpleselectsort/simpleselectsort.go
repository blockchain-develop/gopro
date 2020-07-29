package simpleselectsort

func SimpleSelectSort(data []int32) {
	for i := 0;i < len(data);i ++ {
		key := data[i]
		index := i
		for j := i + 1;j < len(data);j ++ {
			if data[j] < key {
				key = data[j]
				index = j
			}
		}
		data[index] = data[i]
		data[i] = key
	}
}
