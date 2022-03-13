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
		for i := 0; i < (GolHeight * GolWidth); i++ {
			*board.Cell(rand.Intn(GolWidth), rand.Intn(GolHeight)) = true
		}

		for {
			Clear()
			board.Step()
			board.Draw()

			r, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if r[0] == 'n' {
				break
			}
		}

		for j := 0; j < GolHeight; j++ {
			for i := 0; i < GolWidth; i++ {
				*board.Cell(i, j) = false
			}
		}
	}
}
