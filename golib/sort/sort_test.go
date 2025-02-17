package sort

import (
	"github.com/magiconair/properties/assert"
	"math/rand"
	"sort"
	"testing"
)

/*

 */
func TestSort1(t *testing.T) {
	size := 2
	data := make([]int32, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int31()
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func TestSort2(t *testing.T) {
	size := 102400
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}

	which := sort.IntSlice(data)
	sort.Sort(which)
	order := sort.IsSorted(which)
	assert.Equal(t, order, true)
}

type Int32Slice []int32

func (this Int32Slice) Len() int {
	return len(this)
}
func (this Int32Slice) Swap(i, j int) {
	key := this[i]
	this[i] = this[j]
	this[j] = key
}
func (this Int32Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func TestSort3(t *testing.T) {
	size := 102400
	data := make([]int32, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int31()
	}

	sort.Sort(Int32Slice(data))
	order := comsort.IsOrderly(data)
	assert.Equal(t, order, byte(1))
}
