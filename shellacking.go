package main

import (
	"shellacking/shell_core"
	. "shellacking/shell_geo"
	"shellacking/shell_styles"
)

func main() {

	//shell_utils.Clear()
	m := shell_core.CreateEmptyMatrix(40, 100)

	//r := float64(0)

	DrawLine(Point2D{X: 0, Y: 0}, Point2D{X: 30, Y: 10}, shell_styles.BgWhite, &m, &State)
	RotateBy(-10, &State)
	DrawLine(Point2D{X: 0, Y: 0}, Point2D{X: 30, Y: 10}, shell_styles.BgWhite, &m, &State)
	//for true != false {
	//	r -= 10
	//	shell_core.PrintOutputMatrixLn(&m)
	//	time.Sleep(time.Duration(2000) * time.Millisecond)
	//	ResetState()
	//	shell_utils.Clear()
	//}
	//shell_geo.DrawLine(shell_geo.Point2D{X: 20, Y: 39}, shell_geo.Point2D{X: 40, Y: 0}, shell_styles.FgGreen, &m)

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
