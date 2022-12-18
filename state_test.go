package go2048

import (
	"testing"
)


func TestGameState(t *testing.T) {
	gt := NewGameState()
	t.Log("\n" + gt.String())
}
