package shell_geo

// 2x2 . 2x4

func MatrixMult(m [][]float64, t [][]float64) [][]float64 {
	if len(m[0]) != len(t) {
		panic("number of columns of target matrix is different than number of rows of transformation matrix")
	}

	r := make([][]float64, len(m))

	for k := 0; k < len(t[0]); k++ {
		for l := 0; l < len(r); l++ {
			r[l] = *initLine(&r[l], len(t[0]))
			for c := 0; c < len(m[l]); c++ {
				r[l][k] += m[l][c] * t[c][k]
			}
		}
	}

	return r
}

func initLine(line *[]float64, length int) *[]float64 {
	if len(*line) == 0 {
		*line = make([]float64, length)
	}
	return line
}
