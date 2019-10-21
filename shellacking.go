package main

import (
	"fmt"
	"shellacking/shell_styles"
)

type ShellOutput interface {
	format() string
}

func (cs ColoredString) format() string {
	return cs.color.ToString() + cs.text
}

type ColoredString struct {
	text  string
	color shell_styles.Output
}

func createColoredString(text string, color *shell_styles.Output) ShellOutput {
	var outputColoredString ShellOutput = ColoredString{text, *color}
	return outputColoredString
}

type Matrix struct {
	lines [][]ShellOutput
}

func createMatrix(lines [][]ShellOutput) Matrix {
	return Matrix{lines}
}

func printMatrix(matrix *Matrix) {
	for _, row := range matrix.lines {
		var rowLength = len(row)
		for i, value := range row {
			if i == rowLength-1 {
				fmt.Println(value.format())
			} else {
				fmt.Print(value.format())
			}
		}
	}
}

func main() {
	var whatever ShellOutput = createColoredString("whatever", &shell_styles.Blue)
	var iwant ShellOutput = createColoredString("iwant", &shell_styles.Green)
	matrix := createMatrix([][]ShellOutput{{whatever, iwant}, {whatever}})
	printMatrix(&matrix)
}
