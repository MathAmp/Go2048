package go2048

import (
	"testing"
)

func TestFilterZeroIndices(t *testing.T) {
	var ba BlockArray
	ba[0] = 5
	t.Log(FilterZeroIndices(ba))
}

func TestPickNewIndex(t *testing.T) {
	var ba BlockArray

	t.Log(PickNewIndex(ba))
}

func TestMakeDefaultBlock(t *testing.T) {
	ba := make([]Block, 10)

	for i := 0; i < 10; i++ {
		ba[i] = MakeDefaultBlock(0.1)
	}

	t.Log(ba)
}

func TestShiftBlocks(t *testing.T) {
	t.Log(ShiftBlocks([]Block{1, 0, 2, 0, 5}))
}

func TestMergeBlocks(t *testing.T) {
	t.Log(MergeBlocks([]Block{1, 2, 1, 2, 1, 2, 1, 2}))
}

func newArangeBlockArray() BlockArray {
	var ba BlockArray
	for i := 0; i < int(BLOCKS_SIZE*BLOCKS_SIZE); i++ {
		ba[i] = Block(i)
	}
	return ba
}

func testSample(t *testing.T, sampleFunc SampleFunc) {
	ba := newArangeBlockArray()
	sba := sampleFunc(ba)
	t.Log(sba)
	t.Log(sampleFunc(sba))
}

func TestSampleRight(t *testing.T) {
	testSample(t, SampleRight)
}

func TestSampleLeft(t *testing.T) {
	testSample(t, SampleLeft)
}

func TestSampleDown(t *testing.T) {
	testSample(t, SampleDown)
}

func TestSampleUp(t *testing.T) {
	testSample(t, SampleUp)
}
