package activityselection

func ActivieySelection(start []int32, finish []int32) ([]int) {
	selected := make([]int, 0)
	selected = append(selected, 0)
	current_selected := 0
	for i := 1;i < len(finish);i ++ {
		if start[i] >= finish[current_selected] {
			selected = append(selected, i)
			current_selected = i
		}
	}
	return selected
}
