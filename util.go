package Go2048

func GetZeroIndices(bs BlockArray) (r []int) {
	for k, b := range bs {
		if (b == 0) {
			r = append(r, k)
		}
	}
	return r
}
