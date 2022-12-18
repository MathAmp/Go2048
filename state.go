package go2048

import (
	"fmt"
	"time"

	"golang.org/x/exp/slices"
)

type StatusCode int

const (
	TERMINATED StatusCode = iota
	PROCESSED
	NONPROCESSED
)

const BLOCKS_SIZE uint32 = 4

type Block uint32

type BlockArray [BLOCKS_SIZE * BLOCKS_SIZE]Block

func (ba *BlockArray) IterBlocks() func() []Block {
	i := int(0)
	return func() (iba []Block) {
		if i < int(BLOCKS_SIZE) {
			iba = ba[i*int(BLOCKS_SIZE) : (i+1)*int(BLOCKS_SIZE)]
			i++
			return
		}
		iba = nil
		return
	}
}

type GameState struct {
	BlockA    BlockArray
	Ticker    uint32
	StartTime time.Time
}

// Create new Game State
func NewGameState() GameState {
	var blocks BlockArray
	return GameState{blocks, 0, time.Now()}
}

func (gt *GameState) InitRandomBlock() {
	zeroIndices := FilterZeroIndices(gt.BlockA)
	if len(zeroIndices) == 0 {
		panic("Invalid!!")
	}
	gt.BlockA[PickOne(zeroIndices)] = MakeDefaultBlock(0.1)
}

func (gt *GameState) String() string {
	r := fmt.Sprintf("Ticker %d (%s)\n", gt.Ticker, time.Since(gt.StartTime).String())
	for i, v := range gt.BlockA {
		r += fmt.Sprintf("|%6d ", v)

		if uint32(i)%BLOCKS_SIZE == BLOCKS_SIZE-1 {
			r += "|\n"
		}
	}

	return r
}

func (gt *GameState) Process(sd ShiftDirection) StatusCode {
	possibleDirections := GetPossibleDirections(gt.BlockA)
	if len(possibleDirections) == 0 {
		return TERMINATED
	}

	if !slices.Contains(possibleDirections, sd) {
		return NONPROCESSED
	}

	gt.BlockA = ShiftAndMergeBlockArray(gt.BlockA, sd)
	gt.Ticker++
	gt.InitRandomBlock()
	return PROCESSED
}
