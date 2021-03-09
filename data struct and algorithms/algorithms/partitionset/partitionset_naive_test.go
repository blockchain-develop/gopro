package partitionset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionSet_Naive1(t *testing.T) {
	data := [...]int{1, 6, 11, 5}
	set1, set2 := PartitionSet_Naive(data[:])
	assert.Equal(t, len(set1), 1)
	assert.Equal(t, len(set2), 3)
}
