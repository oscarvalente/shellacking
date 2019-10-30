package shell_core

import (
	"fmt"
	"reflect"
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

func (rs ResetString) Format() string {
	return rs.Type.ToString()
}

type ResetString struct {
	Type shell_styles.Output
}

func CreateResetString(reset *shell_styles.Output) ShellOutput {
	var outputResetString ShellOutput = ResetString{*reset}
	return outputResetString
}

type ShellOutput interface {
	Format() string
}

func CloneMatrixLines(lines [][]ShellOutput) [][]ShellOutput {
	clone := [][]ShellOutput{}
	for x, row := range lines {
		clone = append(clone, []ShellOutput{})
		for _, col := range row {
			if reflect.TypeOf(col) == reflect.TypeOf((*ColoredString)(nil)).Elem() {
				clone[x] = append(clone[x], CloneColoredString(col.(ColoredString)))
			} else if reflect.TypeOf(col) == reflect.TypeOf((*ResetString)(nil)).Elem() {
				clone[x] = append(clone[x], CloneResetString(col.(ResetString)))
			}
		}
	}
	return clone
}

func CloneColoredString(cs ColoredString) ColoredString {
	return ColoredString{cs.Text, cs.Color}
}

func CloneResetString(rs ResetString) ResetString {
	return ResetString{rs.Type}
}

func CreateEmptyShellOutputLine() ([]ShellOutput) {
	return []ShellOutput{}
}

type Matrix struct {
	Lines [][]ShellOutput
}

func CreateMatrix(lines [][]ShellOutput) Matrix {
	return Matrix{lines}
}

func CloneMatrix(matrix Matrix) Matrix {
	clone := Matrix{}
	clone.Lines = CloneMatrixLines(matrix.Lines)
	return clone
}

func UpdateMatrixLines(matrix *Matrix, lines [][]ShellOutput) {
	matrix.Lines = lines
}

func PrintMatrix(matrix *Matrix) {
	for _, row := range matrix.Lines {
		for _, value := range row {
			fmt.Print(value.Format())
		}
	}
}

func PrintMatrixLn(matrix *Matrix) {
	matrixText := ""
	for _, row := range matrix.Lines {
		for _, value := range row {
			matrixText += value.Format()
		}
		matrixText += "\n"
	}
	fmt.Print(matrixText)
}

type GradientEffect interface {
	Create(text string, from shell_styles.Output, to shell_styles.Output)
	Play(duration int)
}

type RandomColorsEffect interface {
	Create(iterations int, rows int, cols int, colors []shell_styles.Output)
	Play(duration int)
}

type MatrixreloadedColorsEffect interface {
	Create(chars []rune, rows int, cols int, nSequences int, colors []shell_styles.Output)
	Play(duration int)
}
