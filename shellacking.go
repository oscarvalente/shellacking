package main

import (
	"shellacking/shell_effects"
	"shellacking/shell_styles"
)

func main() {
	myFirstEffect := shell_effects.CreateSequencedLayeredEffect("now as the river dissolves the sea", shell_styles.Blue, shell_styles.Green)
	myFirstEffect.Play(8)
}
