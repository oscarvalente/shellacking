package shell_effects

import (
	"shellacking/shell_core"
	"shellacking/shell_styles"
	"shellacking/shell_utils"
	"time"
)

type SequencedEffect struct {
	sequences []shell_core.Matrix
}

func CreateSequencedLayeredEffect(text string, from shell_styles.Output, to shell_styles.Output) shell_core.LayeredEffect {
	var layeredEffect shell_core.LayeredEffect = &SequencedEffect{[]shell_core.Matrix{}}
	layeredEffect.Create(text, from, to)
	return layeredEffect
}

func (effect *SequencedEffect) Create(text string, from shell_styles.Output, to shell_styles.Output) {
	shell_utils.Clear()
	// create matrices
	runes := []rune(text)
	for i, _ := range runes {
		matrix := shell_core.CreateMatrix(
			[][]shell_core.ShellOutput{
				{shell_core.CreateColoredString(string(runes[:i]), &to)}, {shell_core.CreateColoredString(string(runes[i:]), &from)}})
		effect.sequences = append(effect.sequences, matrix)
	}
}

func (effect SequencedEffect) Play(duration int) {
	//print matrices
	var slice int64
	slice = int64(float64(duration) / float64(len(effect.sequences)) * 1000)

	for i, matrix := range effect.sequences {
		shell_core.PrintMatrix(&matrix)
		if i != len(effect.sequences)-1 {
			time.Sleep(time.Duration(slice) * time.Millisecond)
		}
		shell_utils.Clear()
	}
}