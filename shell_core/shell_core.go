package shell_core

import (
	"fmt"
	"shellacking/shell_styles"
)

func (cs ColoredString) Format() string {
	return cs.Color.ToString() + cs.Text
}

type ColoredString struct {
	Text  string
	Color shell_styles.Output
}

func CreateColoredString(text string, color *shell_styles.Output) ShellOutput {
	var outputColoredString ShellOutput = ColoredString{text, *color}
	return outputColoredString
}

type ShellOutput interface {
	Format() string
}
type Matrix struct {
	Lines [][]ShellOutput
}

func CreateMatrix(lines [][]ShellOutput) Matrix {
	return Matrix{lines}
}

func PrintMatrix(matrix *Matrix) {
	for _, row := range matrix.Lines {
		for _, value := range row {
			fmt.Print(value.Format())
		}
	}
}

func PrintMatrixLn(matrix *Matrix) {
	for _, row := range matrix.Lines {
		var rowLength = len(row)
		for i, value := range row {
			if i == rowLength-1 {
				fmt.Println(value.Format())
			} else {
				fmt.Print(value.Format())
			}
		}
	}
}

type LayeredEffect interface {
	Create(text string, from shell_styles.Output, to shell_styles.Output)
	Play(duration int)
}
