package go2048

import (
	"testing"
)

func TestGameState(t *testing.T) {
	gt := NewGameState()
	t.Log("\n" + gt.String())
}

func TestProcess(t *testing.T) {
	gt := NewGameState()
	gt.InitRandomBlock()
	t.Log(gt.String())
	gt.Process(DOWN)
	t.Log(gt.String())
	gt.Process(RIGHT)
	t.Log(gt.String())
	gt.Process(RIGHT)
	t.Log(gt.String())
}
