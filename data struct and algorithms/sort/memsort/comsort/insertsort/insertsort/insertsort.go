package insertsort

func InsertSort(data []int32) {
	for i := 1;i < len(data);i ++ {
		key := data[i]
		var j int = i - 1
		for ;j >= 0;j -- {
			if data[j] <= key {
				break
			}
			data[j + 1] = data[j]
		}
		data[j + 1] = key
	}
}

