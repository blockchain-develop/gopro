package partitionset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionSet_DP1(t *testing.T) {
	data := [...]int{1, 6, 11, 5}
	diff := PartitionSet_DP(data[:])
	assert.Equal(t, diff, 1)
}
