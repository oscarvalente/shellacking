package main

import (
	"shellacking/shell_core"
	"shellacking/shell_styles"
	. "shellacking/shell_geo"
	"time"
)

func main() {

	m := shell_core.CreateEmptyMatrix(3, 3)
	r := float64(0)

	for r < 18000 {
		TranslateBy(-1, -1, &State)
		RotateBy(r, &State)
		TranslateBy(1, 1, &State)
		CreateLine2D(Point2D{X: 1, Y: 0}, Point2D{X: 1, Y: 2}, shell_styles.BgWhite, &m, &State)
		CreatePoint2D(Point2D{X: 1, Y: 1}, shell_styles.BgGreen, &m, &State)
		shell_core.DrawFrameWithClear(&m)
		r += 45
		ResetState()
		time.Sleep(time.Duration(50) * time.Millisecond)
	}

	time.Sleep(time.Duration(3000) * time.Millisecond)

	//myFirstEffect := basic_gradient_layered.CreateSequencedLayeredEffect("now as the river dissolves the sea", shell_styles.FgBlue, shell_styles.FgGreen)
	//myFirstEffect.Play(2)
	//mySecondEffect := basic_reversed_gradient_layered.CreateBasicReversedLayeredEffect("now as the river dissolves the sea", shell_styles.FgBlue, shell_styles.FgGreen)
	//mySecondEffect.Play(2)
	//myThirdEffect := thering_random_layered.CreateRandomLayeredEffect(50, 20, 160, []shell_styles.Output{shell_styles.BgDefault, shell_styles.BgWhite})
	//myThirdEffect.Play(5)

	//myFourthEffect := matrixreload_random_layered.CreatesRandomLayeredEffect([]rune("日本語"), 10, 80, 80,
	//	[]shell_styles.Output{shell_styles.FgGreen, shell_styles.FgLightGreen, shell_styles.FgDarkGray})
	//myFourthEffect.Play(10)
}
