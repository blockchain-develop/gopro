package uuid

import (
	"github.com/pborman/uuid"
	"testing"
)

func TestUUid(t *testing.T) {
	x := uuid.NewUUID()
	x.Id()
}
