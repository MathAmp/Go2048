package Go2048

import (
	"testing"
)

func TestGetZeroIndices(t *testing.T) {
	var ba BlockArray
	ba[0] = 5
	t.Log(GetZeroIndices(ba))
}