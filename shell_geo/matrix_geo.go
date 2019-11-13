package shell_geo

import (
	"fmt"
	"math"
	"shellacking/shell_core"
	"shellacking/shell_styles"
)

func MatrixMult(m [][]float64, t [][]float64) [][]float64 {
	if len(m[0]) != len(t) {
		panic("number of columns of target matrix is different than number of rows of transformation matrix")
	}

	r := make([][]float64, len(m))

	for k := 0; k < len(t[0]); k++ {
		for l := 0; l < len(r); l++ {
			r[l] = *shell_core.InitMatrixLine(&r[l], len(t[0]))
			for c := 0; c < len(m[l]); c++ {
				r[l][k] += m[l][c] * t[c][k]
			}
		}
	}

	return r
}

func MatrixTranspose(m [][]float64) [][]float64 {
	matrixLines := len(m)
	matrixColumns := len(m[0])
	if matrixLines <= 0 || matrixColumns <= 0 {
		panic("matrix must have rows and columns")
	}

	r := make([][]float64, matrixColumns)

	for c := 0; c < matrixColumns; c++ {
		r[c] = []float64{}
		for l := 0; l < matrixLines; l++ {
			r[c] = append(r[c], m[l][c])
		}
	}

	return r
}

func CreateTranslationMatrix2D(deltaX float64, deltaY float64) [][]float64 {
	// Create Translation T matrix
	// M[l x c] . T[c x c]

	dim := 3
	lastColIndex := dim - 1

	r := make([][]float64, dim)

	for l := 0; l < dim; l++ {
		r[l] = []float64{}
		for c := 0; c < dim; c++ {
			if l == 0 && c == lastColIndex {
				r[l] = append(r[l], deltaX)
			} else if l == 1 && c == lastColIndex {
				r[l] = append(r[l], deltaY)
			} else if l == c {
				r[l] = append(r[l], 1)
			} else {
				r[l] = append(r[l], 0)
			}
		}
	}

	return r
}

func CreateScaleMatrix2D(scaleX float64, scaleY float64) [][]float64 {
	// Create Scale T matrix
	// M[l x c] . T[c x c]

	dim := 3

	r := make([][]float64, dim)

	for l := 0; l < dim; l++ {
		r[l] = []float64{}
		for c := 0; c < dim; c++ {
			if l == 0 && l == c {
				r[l] = append(r[l], scaleX)
			} else if l == 1 && l == c {
				r[l] = append(r[l], scaleY)
			} else if l == c {
				r[l] = append(r[l], 1)
			} else {
				r[l] = append(r[l], 0)
			}
		}
	}

	return r
}

func CreateRotationMatrix2D(angleDegrees float64) [][]float64 {
	// Create Scale T matrix
	// M[l x c] . T[c x c]

	angleRad := angleDegrees * math.Pi / 180

	return [][]float64{{math.Cos(angleRad), -math.Sin(angleRad), 0}, {math.Sin(angleRad), math.Cos(angleRad), 0}, {0, 0, 1}}
}

// using Column Major Form
func TransformPoint2D(p *Point2D, state *[][][]float64) [][]float64 {
	r := [][]float64{{p.X}, {p.Y}, {1}}
	transformations := 0
	for transformations < len(*state) {
		t := (*state)[transformations]
		r = MatrixMult(t, r)
		transformations++
	}

	return r
}

func DrawLine(originalA Point2D, originalB Point2D, color shell_styles.Output, m *shell_core.Matrix, state *[][][]float64) {
	var a, b Point2D
	if (len(*state)) > 0 {
		// transform a
		aTransformMatrix := TransformPoint2D(&originalA, state)
		// transform b
		bTransformMatrix := TransformPoint2D(&originalB, state)
		a = CreatePointFromMatrix2D(aTransformMatrix)
		b = CreatePointFromMatrix2D(bTransformMatrix)
	} else {
		a = originalA
		b = originalB
	}
	yStart := math.Min(a.Y, b.Y)
	yEnd := math.Max(a.Y, b.Y)
	matrixRows := float64(len(m.Lines))
	r := make([][]float64, 2)
	r[0] = []float64{}
	r[1] = []float64{}

	if yStart < 0 || yStart >= matrixRows {
		panic(fmt.Sprintf("starting Y should be within matrix range (start Y: %f)", yStart))
	}

	if yEnd < 0 || yEnd >= matrixRows {
		panic(fmt.Sprintf("ending Y should be within matrix range (end Y: %f)", yEnd))
	}

	linearFuncByY := LinearSlopeFuncByY(a, b)
	for y := yStart; y < yEnd; y++ {
		x := linearFuncByY(y)
		r[0] = append(r[0], x)
		r[1] = append(r[1], y)
		// update by reference
		shell_core.InsertMatrixValue(shell_core.CreateColoredAndResetString(".", &color, &shell_styles.ResetAll), int(x), int(y), m)
	}
}
