package bucketsort

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

type bucket struct {
	items []int32
}

func (this *bucket) insert(data int32) {
	this.items = append(this.items, data)
	key := data
	i := len(this.items) - 2
	for i = len(this.items) - 2; i >= 0;i -- {
		if this.items[i] > data {
			this.items[i + 1] = this.items[i]
		} else {
			break
		}
	}
	this.items[i + 1] = key
}

func BucketSort(data []int32) {
	max, min, size := bound(data)
	buckets := make([]bucket, size)
	base := int32((int64(max - min + 1) + int64(size - 1)) / int64(size))
	for _, item := range data {
		buckets[(item - min) / base].insert(item)
	}

	//
	index := 0
	for i := 0;i < int(size) ;i ++ {
		for _, item := range buckets[i].items {
			data[index] = item
			index ++
		}
	}
}