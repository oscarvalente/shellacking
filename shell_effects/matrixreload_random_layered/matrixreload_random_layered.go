package matrixreload_random_layered

import (
	"shellacking/shell_core"
	"shellacking/shell_styles"
	"shellacking/shell_utils"
	"time"
)

type SequencedEffect struct {
	sequences []shell_core.Matrix
}

func CreatesRandomLayeredEffect(chars []rune, rows int, cols int, nSequences int, colors []shell_styles.Output) shell_core.MatrixreloadedColorsEffect {
	var layeredEffect shell_core.MatrixreloadedColorsEffect = &SequencedEffect{[]shell_core.Matrix{}}
	layeredEffect.Create(chars, rows, cols, nSequences, colors)
	return layeredEffect
}

func (effect *SequencedEffect) Create(chars []rune, rows int, cols int, nSequences int, colors []shell_styles.Output) {
	shell_utils.Clear()
	// create 1st matrix
	lines := [][]shell_core.ShellOutput{{}}
	for x := 0; x < rows; x++ { // lines
		lines = append(lines, shell_core.CreateEmptyShellOutputLine())
		for y := 0; y < cols; y++ { //cols
			randomColorIndex := shell_utils.RandInt(0, len(colors)-1)
			randomCharIndex := shell_utils.RandInt(0, len(chars)-1)
			coloredString := shell_core.CreateColoredString(string(chars[randomCharIndex]), &colors[randomColorIndex])
			lines[x] = append(lines[x], coloredString)
			if x == rows-1 && y == cols-1 {
				resetString := shell_core.CreateResetString(&shell_styles.ResetAll)
				lines[x] = append(lines[x], resetString)
			}
		}
	}

	matrix := shell_core.CreateMatrix(lines)
	for t := 0; t < nSequences; t++ {
		if t > 0 {
			matrix = shell_core.CloneMatrix(effect.sequences[t-1])
			//fmt.Print(matrix)
			shell_core.UpdateMatrixLines(&matrix, populateRemainingLines(matrix.Lines, chars, colors, cols))
		}
		effect.sequences = append(effect.sequences, matrix)
	}
}

func (effect SequencedEffect) Play(duration int) {
	//print matrices
	var slice int64
	slice = int64(float64(duration) / float64(len(effect.sequences)) * 1000)

	for i, matrix := range effect.sequences {
		shell_core.PrintOutputMatrixLn(&matrix)
		if i != len(effect.sequences)-1 {
			time.Sleep(time.Duration(slice) * time.Millisecond)
		}
		shell_utils.Clear()
	}
}

func populateRemainingLines(lines [][]shell_core.ShellOutput, chars []rune, colors []shell_styles.Output, cols int) [][]shell_core.ShellOutput {
	for y := 0; y < cols; y++ { //cols
		randomColorIndex := shell_utils.RandInt(0, len(colors)-1)
		randomCharsIndex := shell_utils.RandInt(0, len(chars)-1)
		coloredString := shell_core.CreateColoredString(string(chars[randomCharsIndex]), &colors[randomColorIndex])
		lines[0][y] = coloredString
	}
	return lines
}
