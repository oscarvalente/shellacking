package shell_geo

import (
	"math"
	"shellacking/shell_core"
	"shellacking/shell_styles"
)

type Point struct {
	X float64
	Y float64
}

func DrawLine(a Point, b Point, color shell_styles.Output, m *shell_core.Matrix) {
	yStart := math.Min(a.Y, b.Y)
	yEnd := math.Max(a.Y, b.Y)
	matrixRows := float64(len(m.Lines))
	if yStart < 0 || yStart >= matrixRows {
		panic("starting Y should be within matrix range")
	}

	if yEnd < 0 || yEnd >= matrixRows {
		panic("starting Y should be within matrix range")
	}
	linearFuncByY := LinearSlopeFuncByY(a, b)
	for y := yStart; y < yEnd; y++ {
		x := linearFuncByY(y)
		// update by reference
		shell_core.InsertMatrixValue(shell_core.CreateColoredAndResetString(".", &color, &shell_styles.ResetAll), int(x), int(y), m)
	}
}

func AngleOfSlopeBetween(a Point, b Point) float64 {
	distance := distanceBetween(a, b)
	angle := math.Acos(deltaX(a, b)/distance) * 100
	return angle
}

func linearSlope(a Point, b Point) float64 {
	return deltaY(a, b) / deltaX(a, b)
}

func linearYX0(a Point, b Point) float64 {
	slope := linearSlope(a, b)
	return a.Y - slope*a.X
}

func LinearSlopeFuncByX(a Point, b Point) func(x float64) float64 {
	slope := linearSlope(a, b)
	yx0 := linearYX0(a, b)
	return func(x float64) float64 {
		return x*slope + yx0
	}
}

func LinearSlopeFuncByY(a Point, b Point) func(x float64) float64 {
	slope := linearSlope(a, b)
	yx0 := linearYX0(a, b)
	return func(y float64) float64 {
		return (y - yx0) / slope
	}
}

func distanceBetween(a Point, b Point) float64 {
	return math.Sqrt(math.Pow(deltaX(a, b), 2) + math.Pow(deltaY(a, b), 2))
}

func deltaX(a Point, b Point) float64 {
	return b.X - a.X
}

func deltaY(a Point, b Point) float64 {
	return b.Y - a.Y
}
