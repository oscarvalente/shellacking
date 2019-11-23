package shell_geo

import (
	"errors"
	"math"
)

type Point2D struct {
	X float64
	Y float64
}

// create point from a matrix in Column Major Form
func CreatePointFromMatrix2D(m [][]float64) Point2D {
	return Point2D{X: m[0][0], Y: m[1][0]}
}

func AngleOfSlopeBetween(a Point2D, b Point2D) float64 {
	distance := distanceBetween(a, b)
	angle := math.Acos(deltaX(a, b)/distance) * 100
	return angle
}

func linearSlope(a Point2D, b Point2D) (float64, error) {
	deltaX := deltaX(a, b)
	if (deltaX == 0) {
		return 0, errors.New("Slope calculation: cannot divide by 0")
	}
	return deltaY(a, b) / deltaX, nil
}

func linearYX0(a Point2D, b Point2D) float64 {
	slope, _ := linearSlope(a, b)
	return a.Y - slope*a.X
}

func LinearSlopeFuncByX(a Point2D, b Point2D) func(x float64) float64 {
	slope, _ := linearSlope(a, b)
	yx0 := linearYX0(a, b)
	return func(x float64) float64 {
		return x*slope + yx0
	}
}

func LinearSlopeFuncByY(a Point2D, b Point2D) func(x float64) float64 {
	slope, err := linearSlope(a, b)

	yx0 := linearYX0(a, b)
	return func(y float64) float64 {
		if err == nil {
			return (y - yx0) / slope
		} else {
			return a.X
		}
	}
}

func distanceBetween(a Point2D, b Point2D) float64 {
	return math.Sqrt(math.Pow(deltaX(a, b), 2) + math.Pow(deltaY(a, b), 2))
}

func deltaX(a Point2D, b Point2D) float64 {
	return b.X - a.X
}

func deltaY(a Point2D, b Point2D) float64 {
	return b.Y - a.Y
}
