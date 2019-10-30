package thering_random_layered

import (
	"shellacking/shell_core"
	"shellacking/shell_styles"
	"shellacking/shell_utils"
	"time"
)

type SequencedEffect struct {
	sequences []shell_core.Matrix
}

func CreateRandomLayeredEffect(iterations int, rows int, cols int, colors []shell_styles.Output) shell_core.RandomColorsEffect {
	var layeredEffect shell_core.RandomColorsEffect = &SequencedEffect{[]shell_core.Matrix{}}
	layeredEffect.Create(iterations, rows, cols, colors)
	return layeredEffect
}

func (effect *SequencedEffect) Create(iterations int, rows int, cols int, colors []shell_styles.Output) {
	shell_utils.Clear()
	// create matrices
	for t := 0; t < iterations; t++ {
		lines := [][]shell_core.ShellOutput{{}}
		for x := 0; x < rows; x++ { // lines
			lines = append(lines, shell_core.CreateEmptyShellOutputLine())
			for y := 0; y < cols; y++ { //cols
				randomColorIndex := shell_utils.RandInt(0, len(colors)-1)
				coloredString := shell_core.CreateColoredString(" ", &colors[randomColorIndex])
				lines[x] = append(lines[x], coloredString)
				if x == rows-1 && y == cols-1 {
					resetString := shell_core.CreateResetString(&shell_styles.ResetAll)
					lines[x] = append(lines[x], resetString)
				}
			}
		}
		matrix := shell_core.CreateMatrix(lines)
		effect.sequences = append(effect.sequences, matrix)
	}
}

func (effect SequencedEffect) Play(duration int) {
	//print matrices
	var slice int64
	slice = int64(float64(duration) / float64(len(effect.sequences)) * 1000)

	for i, matrix := range effect.sequences {
		shell_core.PrintMatrixLn(&matrix)
		if i != len(effect.sequences)-1 {
			time.Sleep(time.Duration(slice) * time.Millisecond)
		}
		shell_utils.Clear()
	}
}
