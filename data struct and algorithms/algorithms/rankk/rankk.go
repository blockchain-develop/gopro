package rankk

func random_partition(data []int32, p int, q int) int {
	if p == q {
		return p
	}

	key := data[p]
	d := 1
	for p < q {
		if d == 1 {
			if data[q] >= key {
				q --
			} else {
				data[p] = data[q]
				p ++
				d = 0
			}
		} else {
			if data[p] <= key {
				p ++
			} else {
				data[q] = data[p]
				q --
				d = 1
			}
		}
	}
	data[p] = key
	return p
}

func random_select(data []int32, p int, q int, k int) int32 {
	r := random_partition(data, p, q)
	i := r - p + 1
	if i == k {
		return data[r]
	} else if i > k {
		return random_select(data, p, r - 1, k)
	} else {
		return random_select(data, r + 1, q, k - i)
	}
}

func RankK(data []int32, k int) int32 {
	return random_select(data, 0, len(data) - 1, k)
}