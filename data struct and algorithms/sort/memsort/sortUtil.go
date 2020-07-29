package comsort

/*
if unorder, return 0
if from smaller to biger, return 1
if from biger to smaller, return 2
 */
func IsOrderly(data []int32) (byte) {
	if len(data) == 0 || len(data) == 1 {
		return 1
	}
	var order byte = 0
	temp := data[0]
	for i := 1;i < len(data);i ++ {
		if data[i] < temp {
			if order == 0 {
				order = 2
			} else if order == 1 {
				return 0
			} else {

			}
			temp = data[i]
		} else if data[i] > temp {
			if order == 0 {
				order = 1
			} else if order == 2 {
				return 0
			} else {

			}
			temp = data[i]
		} else {

		}
	}
	if order == 0 {
		return 1
	}
	return order
}
