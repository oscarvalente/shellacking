package shell_core

import (
	"fmt"
	"reflect"
	"shellacking/shell_styles"
	"shellacking/shell_utils"
)

func (cs ColoredString) Format() string {
	return cs.Color.ToString() + cs.Text
}

func (cs ColoredString) isEmpty() bool {
	return cs.Text == ""
}

type ColoredString struct {
	Text  string
	Color shell_styles.Output
}

func CreateColoredString(text string, color *shell_styles.Output) ShellOutput {
	var outputColoredString ShellOutput = ColoredString{text, *color}
	return outputColoredString
}

func (coloredResetString ColoredAndResetString) Format() string {
	return coloredResetString.Color.ToString() + coloredResetString.Text + coloredResetString.Type.ToString()
}

func (coloredResetString ColoredAndResetString) isEmpty() bool {
	return coloredResetString.Text == ""
}

type ColoredAndResetString struct {
	Text  string
	Color shell_styles.Output
	Type  shell_styles.Output
}

func CreateColoredAndResetString(text string, color *shell_styles.Output, reset *shell_styles.Output) ShellOutput {
	var outputColoredString ShellOutput = ColoredAndResetString{text, *color, *reset}
	return outputColoredString
}

func CreateEmptyString() ShellOutput {
	var outputColoredString ShellOutput = ColoredString{" ", shell_styles.FgDefault}
	return outputColoredString
}

func (rs ResetString) Format() string {
	return rs.Type.ToString()
}

func (rs ResetString) isEmpty() bool {
	return true
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
	isEmpty() bool
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

func CreateEmptyMatrix(rows int, cols int) Matrix {
	lines := [][]ShellOutput{{}}
	for x := 0; x < rows; x++ { // lines
		lines = append(lines, CreateEmptyShellOutputLine())
		for y := 0; y < cols; y++ { //cols
			emptyString := CreateEmptyString()
			lines[x] = append(lines[x], emptyString)
		}
	}
	return CreateMatrix(lines)
}

func CloneMatrix(matrix Matrix) Matrix {
	clone := Matrix{}
	clone.Lines = CloneMatrixLines(matrix.Lines)
	return clone
}

func UpdateMatrixLines(matrix *Matrix, lines [][]ShellOutput) {
	matrix.Lines = lines
}

func InsertMatrixValue(value ShellOutput, x int, y int, matrix *Matrix) {
	if y < 0 || y >= len(matrix.Lines) || x < 0 || x >= len(matrix.Lines[0]) {
		panic(fmt.Sprintf("Attempting to insert value outside shell output Matrix range (x: %d, y: %d)", x, y))
	}
	matrix.Lines[y][x] = value
}

func PrintMatrix(matrix *Matrix) {
	for _, row := range matrix.Lines {
		for _, value := range row {
			fmt.Print(value.Format())
		}
	}
}

func PrintOutputMatrixLn(matrix *Matrix) {
	matrixText := ""
	for _, row := range matrix.Lines {
		for _, value := range row {
			matrixText += value.Format()
		}
		matrixText += "\n"
	}
	fmt.Print(matrixText)
}

func PrintMatrixLn(matrix *[][]float64) {
	matrixText := ""
	for x, row := range *matrix {
		for y, _ := range row {
			matrixText += shell_utils.ParseFloat64ToString((*matrix)[x][y])
		}
		if x < len(*matrix)-1 {
			matrixText += ""
		}
	}
	fmt.Print(matrixText)
}

func InitMatrixLine(line *[]float64, length int) *[]float64 {
	if len(*line) == 0 {
		*line = make([]float64, length)
	}
	return line
}

func DrawFrameWithClear(m *Matrix) {
	shell_utils.Clear()
	PrintOutputMatrixLn(m)
	*m = CreateEmptyMatrix(len((*m).Lines) - 1, len((*m).Lines[0]))
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
