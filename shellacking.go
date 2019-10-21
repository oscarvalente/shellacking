package main

import (
	"fmt"
	"shellacking/shell_styles"
	"shellacking/shell_utils"

	//"shellacking/shell_utils"
	"time"
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

func printMatrixLn(matrix *Matrix) {
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

func printMatrix(matrix *Matrix) {
	for _, row := range matrix.lines {
		for _, value := range row {
			fmt.Print(value.format())
		}
	}
}

func createCharGradientEffect(text string, from shell_styles.Output, to shell_styles.Output, duration int) {
	// create matrixes
	runes := []rune(text)
	listOfCombinations := []Matrix{}
	for i, _ := range runes {
		matrix := createMatrix([][]ShellOutput{
			{createColoredString(string(runes[:i]), &to)}, {createColoredString(string(runes[i:]), &from)}})
		listOfCombinations = append(listOfCombinations, matrix)
	}

	//print Matrixes
	var slice int64
	slice = int64(float64(duration)/float64(len(listOfCombinations)) * 1000)

	for _, matrix := range listOfCombinations {
		printMatrix(&matrix)
		time.Sleep(time.Duration(slice) * time.Millisecond)
		shell_utils.Clear()
	}
}

// TODO: create effects entities ^^^^^

func main() {
	//var whatever ShellOutput = createColoredString("whatever", &shell_styles.Blue)
	//var iwant ShellOutput = createColoredString("iwant", &shell_styles.Green)
	//matrix := createMatrix([][]ShellOutput{{whatever, iwant}, {whatever}})
	//printMatrixLn(&matrix)
	createCharGradientEffect("now as the river dissolves the sea", shell_styles.Blue, shell_styles.Green, 3)
}
