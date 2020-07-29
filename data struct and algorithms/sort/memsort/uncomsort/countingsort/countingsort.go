package countingsort

import (
	"math"
)

func bound(data []int32) (max int32, min int32, size int32) {
	min = math.MaxInt32
	max = math.MinInt32
	size = 0
	for _,item := range data {
		size ++
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}
	return
}
func CountingSort(data []int32) {
	max, min, size := bound(data)
	counter := make([]int, max - min + 1)
	for _, item := range data {
		counter[item - min] ++
	}
	for i := 1;i < int(max - min + 1);i ++ {
		counter[i] = counter[i] + counter[i - 1]
	}
	data1 := make([]int32, size)
	for i := size - 1;i >= 0;i -- {
		counter[(data[i] - min)] = counter[(data[i] - min)] - 1
		data1[counter[(data[i] - min)]] = data[i]
	}
	for i := 0;i < int(size);i ++ {
		data[i] = data1[i]
	}
}