package maximumsubarray

import "math"

func MaximumSubarray_Force(data []int32) (int, int, int) {
	start,end := 0, 0
	max := 0
	for i := 0;i < len(data);i ++ {
		for j := len(data) - 1;j >= i;j -- {
			sum := 0
			for k := i;k <= j;k ++ {
				sum += int(data[k])
			}
			if sum > max {
				max = sum
				start,end = i,j
			}
		}
	}
	return max, start, end
}
func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func MaximumSubarray_Kadane(data []int) int {
	max_end_here, max_so_far := math.MinInt32, math.MinInt32
	for _, item := range data {
		max_end_here = max(max_end_here + item, item)
		max_so_far = max(max_so_far, max_end_here)
	}
	return max_so_far
}

func MaximumSubarray_KadaneWithIndex(data []int) (int, int, int) {
	max_end_here, max_so_far := math.MinInt32, math.MinInt32
	max_end_here_front, max_so_far_front, max_so_far_end := 0, 0, 0
	for i, item := range data {
		if max_end_here + item > item {
			max_end_here += item
		} else {
			max_end_here = item
			max_end_here_front = i
		}
		if max_so_far > max_end_here {
		} else {
			max_so_far = max_end_here
			max_so_far_front = max_end_here_front
			max_so_far_end = i
		}
	}
	return max_so_far, max_so_far_front, max_so_far_end
}
