package Go2048

import (
	"math/rand"
)

func FilterZeroIndices(bs BlockArray) (r []int) {
	for k, b := range bs {
		if b == 0 {
			r = append(r, k)
		}
	}
	return r
}

func PickNewIndex(bs BlockArray) int {
	zis := FilterZeroIndices(bs)
	return rand.Intn(len(zis))
}

func MakeDefaultBlock(probFour float64) Block {
	if rand.Float64() < probFour {
		return Block(4)
	}

	return Block(2)
}

func fillZeros(bs []Block, total int) []Block {
	requireCount := total - len(bs)
	if requireCount <= 0 {
		return bs
	}

	return append(bs, make([]Block, requireCount)...)[:total]
}

// New block slice, isChanged
func ShiftBlocks(bs []Block) (_ []Block, isChanged bool) {
	nbs := make([]Block, 0)
	isZeroDetected := false

	for _, v := range bs {
		if v != 0 {
			nbs = append(nbs, v)
			isChanged = isZeroDetected
			continue
		}
		isZeroDetected = true
	}

	nbs = fillZeros(nbs, len(bs))
	return nbs, isChanged
}

func MergeBlocks(bs []Block) (_ []Block, isChanged bool) {
	prev := Block(0)
	nbs := make([]Block, 0)

	for _, v := range bs {
		if v == Block(0) {
			break
		}

		if prev == Block(0) {
			prev = v
			continue
		}

		if prev == v {
			nbs = append(nbs, prev+v)
			prev = Block(0)
			isChanged = true
			continue
		}

		nbs = append(nbs, prev)
		prev = v
	}

	nbs = append(nbs, prev)

	return fillZeros(nbs, len(bs)), isChanged
}

type SampleFunc func(BlockArray) BlockArray

func SampleRight(ba BlockArray) BlockArray {
	return ba
}

func SampleLeft(ba BlockArray) (nba BlockArray) {
	for i, v := range ba {
		x := uint32(i) % BLOCKS_SIZE
		y := uint32(i) / BLOCKS_SIZE
		nba[y*BLOCKS_SIZE+(BLOCKS_SIZE-x-1)] = v
	}
	return
}

func SampleDown(ba BlockArray) (nba BlockArray) {
	for i, v := range ba {
		x := uint32(i) % BLOCKS_SIZE
		y := uint32(i) / BLOCKS_SIZE
		nba[x*BLOCKS_SIZE+y] = v
	}
	return
}

func SampleUp(ba BlockArray) (nba BlockArray) {
	for i, v := range ba {
		x := uint32(i) % BLOCKS_SIZE
		y := uint32(i) / BLOCKS_SIZE
		nba[(BLOCKS_SIZE-x-1)*BLOCKS_SIZE+(BLOCKS_SIZE-y-1)] = v
	}
	return
}
