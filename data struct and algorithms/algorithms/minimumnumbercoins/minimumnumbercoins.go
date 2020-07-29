package minimumnumbercoins

func MinimumNumberCoins(value int, vrs []int) (coins []int) {
	currency := value
	coins = make([]int, 0)
	for i := len(vrs) - 1;i >= 0; {
		if currency >= vrs[i] {
			coins = append(coins, vrs[i])
			currency -= vrs[i]
		} else {
			i --
		}
	}
	return
}
