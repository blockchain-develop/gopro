package rankk

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestRankK1(t *testing.T) {
	data := [...]int32{1, 2, 3, 4, 5}
	k := RankK(data[:], 3)
	assert.Equal(t, k, int32(3))
}

func TestRankK2(t *testing.T) {
	data := [...]int32{1, 2, 3, 4}
	k := RankK(data[:], 1)
	assert.Equal(t, k, int32(1))
}

func TestRankK3(t *testing.T) {
	data := [...]int32{1, 2, 2, 4}
	k := RankK(data[:], 4)
	assert.Equal(t, k, int32(4))
}

func TestRankK4(t *testing.T) {
	data := [...]int32{1, 2, 2, 2, 4}
	k := RankK(data[:], 2)
	assert.Equal(t, k, int32(2))
}

func TestRankK5(t *testing.T) {
	data := [...]int32{1, 2}
	k := RankK(data[:], 1)
	assert.Equal(t, k, int32(1))
}

func TestRankK6(t *testing.T) {
	data := [...]int32{2, 2}
	k := RankK(data[:], 2)
	assert.Equal(t, k, int32(2))
}

func TestRankK7(t *testing.T) {
	size := 10240000
	data := make([]int32, size)
	data1 := make([]int32, size)
	for i := 0;i < size;i ++ {
		data[i] = rand.Int31()
		data1[i] = data[i]
	}
	k := RankK(data[:], 10)

	// sort
	sort.Slice(data1, func(i, j int) bool {
		return data1[i] < data1[j]
	})

	// middle
	sortk := data1[9]

	//
	assert.Equal(t, sortk, k)
}
