package main

import (
	"bufio"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

const platform = runtime.GOOS

// multi platform console clear
func Clear() {
	var cmd *exec.Cmd
	if platform == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if platform == "linux" {
		cmd = exec.Command("clear")
	} else {
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	board := NewBoard()
	reader := bufio.NewReader(os.Stdin)

	for {
		rand.Seed(rand.Int63())
		for i := 0; i < rand.Intn(20); i++ {
			*board.Cell(rand.Intn(GolWidth), rand.Intn(GolHeight)) = true
		}

		for {
			Clear()
			board.Step()
			board.Draw()

			r, _, err := reader.ReadRune()
			if err != nil {
				panic(err)
			}
			if r == 'n' {
				break
			}
		}
	}
}
