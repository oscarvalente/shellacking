package shell_utils

import (
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
