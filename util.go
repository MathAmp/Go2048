package go2048

import (
	"math/rand"
)

type ShiftDirection int

const (
	UP ShiftDirection = iota
	DOWN
	LEFT
	RIGHT
)

func FilterZeroIndices(bs BlockArray) (r []int) {
	for k, b := range bs {
		if b == 0 {
			r = append(r, k)
		}
	}
	return r
}

func PickOne(indices []int) int {
	return indices[rand.Intn(len(indices))]
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

func ShiftAndMergeBlocks(bs []Block) ([]Block, bool) {
	bs, isShifted := ShiftBlocks(bs)
	bs, isMerged := MergeBlocks(bs)

	return bs, isShifted || isMerged
}

func ShiftAndMergeBlockArray(ba BlockArray, sd ShiftDirection) (nba BlockArray) {
	sampleFunc := ShiftDirectionToSampleFunc(sd)

	ba = sampleFunc(ba)
	iba := ba.IterBlocks()
	inba := nba.IterBlocks()
	for {
		vba := iba()
		vnba := inba()
		if vba == nil || vnba == nil {
			nba = sampleFunc(nba)
			return
		}

		vba, _ = ShiftAndMergeBlocks(vba)
		copy(vnba, vba)
	}
}

func IsPossibleToShiftAndMerge(ba BlockArray, sd ShiftDirection) (isPossible bool) {
	sampleFunc := ShiftDirectionToSampleFunc(sd)

	ba = sampleFunc(ba)
	iba := ba.IterBlocks()

	for {
		v := iba()
		if v == nil {
			return isPossible
		}
		_, isChanged := ShiftAndMergeBlocks(v)
		isPossible = isPossible || isChanged
	}
}

func GetPossibleDirections(ba BlockArray) (directions []ShiftDirection) {
	for _, d := range []ShiftDirection{UP, DOWN, LEFT, RIGHT} {
		if IsPossibleToShiftAndMerge(ba, d) {
			directions = append(directions, d)
		}
	}
	return
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

func ShiftDirectionToSampleFunc(sd ShiftDirection) SampleFunc {
	switch sd {
	case UP:
		return SampleDown
	case DOWN:
		return SampleUp
	case RIGHT:
		return SampleLeft
	case LEFT:
		return SampleRight
	default:
		panic("Invalid Direction")
	}
}
