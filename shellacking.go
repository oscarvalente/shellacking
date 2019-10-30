package main

import (
	"shellacking/shell_effects/matrixreload_random_layered"
	"shellacking/shell_styles"
)

func main() {

	//myFirstEffect := basic_gradient_layered.CreateSequencedLayeredEffect("now as the river dissolves the sea", shell_styles.FgBlue, shell_styles.FgGreen)
	//myFirstEffect.Play(2)
	//mySecondEffect := basic_reversed_gradient_layered.CreateBasicReversedLayeredEffect("now as the river dissolves the sea", shell_styles.FgBlue, shell_styles.FgGreen)
	//mySecondEffect.Play(2)
	//myThirdEffect := thering_random_layered.CreateRandomLayeredEffect(50, 20, 160, []shell_styles.Output{shell_styles.BgDefault, shell_styles.BgWhite})
	//myThirdEffect.Play(5)

	myFourthEffect := matrixreload_random_layered.CreatesRandomLayeredEffect([]rune("日本語"), 10, 80, 80,
		[]shell_styles.Output{shell_styles.FgGreen, shell_styles.FgLightGreen, shell_styles.FgDarkGray})
	myFourthEffect.Play(10)
}
