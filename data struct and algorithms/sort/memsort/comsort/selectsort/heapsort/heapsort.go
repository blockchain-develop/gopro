package heapsort

func adjustdown(data []int32, root int, end int) {
	parent := root
	child := parent * 2 + 1
	for child <= end {
		if child + 1 <= end && data[child] < data[child + 1] {
			child = child + 1
		}
		if data[child] <= data[parent] {
			break
		}

		key := data[parent]
		data[parent] = data[child]
		data[child] = key

		parent = child
		child = parent * 2 + 1
	}
}

func HeapSort(data []int32) {
	for i := len(data) / 2 - 1;i >= 0;i -- {
		adjustdown(data, i, len(data) - 1)
	}

	for i := len(data) - 1;i > 0;i -- {
		key := data[i]
		data[i] = data[0]
		data[0] = key

		adjustdown(data, 0, i - 1)
	}
}
