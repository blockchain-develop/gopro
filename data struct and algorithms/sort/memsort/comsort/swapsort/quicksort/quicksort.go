package quicksort

func QuickSort(data []int32) {
	quicksort(data, 0, len(data) - 1)
}

func quicksort(data []int32, start int, end int) {
	if len(data) == 1 {
		return
	}

	key := data[start]
	low := start
	high := end
	d := 1
	for low < high {
		if d == 1 {
			if data[high] < key {
				data[low] = data[high]
				low ++
				d = 0
			} else {
				high --
			}
		} else {
			if data[low] > key {
				data[high] = data[low]
				high --
				d = 1
			} else {
				low ++
			}
		}
	}
	data[low] = key
	if low > start + 1 {
		quicksort(data, start, low-1)
	}
	if high < end - 1 {
		quicksort(data, high+1, end)
	}
}
