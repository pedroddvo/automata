package main

import "fmt"

const (
	GolHeight = 50
	GolWidth  = 100
)

type Cell bool

type Board struct {
	cells [GolHeight * GolWidth]Cell
}

func (b *Board) Cell(x, y int) *Cell {
	return &b.cells[(y*GolWidth)+x]
}

// includes itself
func (b *Board) Neighbours(x, y int) (n int) {
	// constrain size of square for no index error
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			wrappedX := (x + i + GolWidth) % GolWidth
			wrappedY := (y + j + GolHeight) % GolHeight
			if (j != 0 || i != 0) && *b.Cell(wrappedX, wrappedY) {
				n++
			}
		}
	}

	return
}

func (b *Board) Step() {
	// copy board
	copy := *b
	for j := 0; j < GolHeight; j++ {
		for i := 0; i < GolWidth; i++ {
			n, cell := b.Neighbours(i, j), *b.Cell(i, j)
			stepCell := (!cell && n == 3) || (cell && (n == 3 || n == 2))
			*copy.Cell(i, j) = stepCell
		}
	}

	*b = copy
}

func (b *Board) Draw() {
	for j := 0; j < GolHeight; j++ {
		for i := 0; i < GolWidth; i++ {
			if *b.Cell(i, j) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func NewBoard() Board {
	return Board{}
}
