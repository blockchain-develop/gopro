package radixsort

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

func getbits(item int32) int {
	bit := 0
	for item > 0 {
		bit ++
		item = item / 10
	}
	return bit
}

func radixsort(data[]int32, data1[]int32, bit int) {
	counter := make([]int, 10)
	base := 1
	for i := 0;i < bit;i ++ {
		base *= 10
	}
	for _, item := range data {
		x := item % (int32(base))
		x = x / (int32(base/10))
		counter[x] ++
	}
	for i := 1;i < 10;i ++ {
		counter[i] = counter[i] + counter[i - 1]
	}
	for i := len(data) - 1;i >= 0;i -- {
		x := data[i] % (int32(base))
		x = x / (int32(base/10))
		counter[x] --
		data1[counter[x]] = data[i]
	}
}

func RadixSort(data []int32) {
	max, _, size := bound(data)
	data1 := make([]int32, len(data))
	bits := getbits(max)
	for i := 0;i < bits;i ++ {
		radixsort(data, data1, i + 1)
		for j := 0;j < int(size);j ++ {
			data[j] = data1[j]
		}
	}
}