package shell_utils

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max+1-min)
}

func ParseFloat64ToString(float float64) string {
	return fmt.Sprintf("%f", float)
}

func DegreesToRad(degrees float64) float64 {
	return degrees * math.Pi / 180
}
