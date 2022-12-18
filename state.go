package go2048

import (
	"fmt"
	"time"
)

const BLOCKS_SIZE uint32 = 4

type Block uint32

type BlockArray [BLOCKS_SIZE * BLOCKS_SIZE]Block

type GameState struct {
	Blocks    BlockArray
	Ticker    uint32
	StartTime time.Time
}

// Create new Game State
func NewGameState() GameState {
	var blocks BlockArray
	return GameState{blocks, 0, time.Now()}
}

func (gt *GameState) String() string {
	r := fmt.Sprintf("Ticker %d (%s)\n", gt.Ticker, time.Since(gt.StartTime).String())
	for i, v := range gt.Blocks {
		r += fmt.Sprintf("|%6d ", v)

		if uint32(i)%BLOCKS_SIZE == BLOCKS_SIZE-1 {
			r += "|\n"
		}
	}

	return r
}
