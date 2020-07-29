package partitionset

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPartitionSet_Recursive1(t *testing.T) {
	data := [...]int{1, 6, 11, 5}
	diff := PartitionSet_Recursive(data[:])
	assert.Equal(t, diff, 1)
}
