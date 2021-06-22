package xxxx


func Problem(num1 []int, num2 []int) {
	problem(num1, num2, 0, len(num1) + len(num2) - 1)
}

func getKey(num1 []int, num2 []int, index int) int {
	if index > len(num1) {
		return num2[index - len(num1)]
	} else {
		return num1[index]
	}
}

func setKey(num1 []int, num2 []int, index int, data int) {
	if index > len(num1) {
		num2[index - len(num1)] = data
	} else {
		num1[index] = data
	}
}

func problem(num1 []int, num2 []int, start int, end int) {
	key := getKey(num1, num2, start)
	low := start
	high := end
	d := 1
	for low < high {
		if d == 1 {
			if getKey(num1, num2, high) < key {
				setKey(num1, num2, low, getKey(num1, num2, start))
				low ++
				d = 0
			} else {
				high --
			}
		} else {
			if getKey(num1, num2, low) > key {
				setKey(num1, num2, high, getKey(num1, num2, low))
				high --
				d = 1
			} else {
				low ++
			}
		}
	}
	if low > start + 1 {
		problem(num1, num2, start, low-1)
	}
	if high < end - 1 {
		problem(num1, num2, high+1, end)
	}
}