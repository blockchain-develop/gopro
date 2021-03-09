package activityselection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActivieySelection1(t *testing.T) {
	start := []int32{10, 12, 20}
	finish := []int32{20, 25, 30}
	selected := ActivieySelection(start, finish)

	assert.Equal(t, len(selected), 2)
	assert.Equal(t, selected, []int{0, 2})
}
func TestActivieySelection2(t *testing.T) {
	start := []int32{1, 3, 0, 5, 8, 5}
	finish := []int32{2, 4, 6, 7, 9, 9}
	selected := ActivieySelection(start, finish)

	assert.Equal(t, len(selected), 4)
	assert.Equal(t, selected, []int{0, 1, 3, 4})
}

