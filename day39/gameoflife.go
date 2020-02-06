package day39

/*
Conway's Game of Life takes place on an infinite two-dimensional board of square cells. Each cell is either dead or alive, and at each tick, the following rules apply:

Any live cell with less than two live neighbours dies.
Any live cell with two or three live neighbours remains living.
Any live cell with more than three live neighbours dies.
Any dead cell with exactly three live neighbours becomes a live cell.
A cell neighbours another cell if it is horizontally, vertically, or diagonally adjacent.

Implement Conway's Game of Life. It should be able to be initialized with a starting list of live cell coordinates and the number of steps it should run for. Once initialized, it should print out the board state at each step. Since it's an infinite board, print out only the relevant coordinates, i.e. from the top-leftmost live cell to bottom-rightmost live cell.

You can represent a live cell with an asterisk (*) and a dead cell with a dot (.).
*/

import "fmt"

// Cell represents single cell on a board
type Cell struct {
	Row int
	Col int
}

// Board represents visible part of inifinite board
type Board struct {
	Alive       map[int]map[int]interface{}
	topLeft     Cell
	bottomRight Cell
}

// Solve plays Game of Life and returns all states of a board
func Solve(turns int, seed Cell, rest ...Cell) []Board {
	boards := []Board{initBoard(append(rest, seed))}
	for i := 0; i < turns; i++ {
		boards := append(boards, nextTurn(boards[i-1]))
	}
	return boards
}

// Print prints out a board state
func (b Board) Print() {
	for i := b.topLeft.Row; i <= b.bottomRight.Row; i++ {
		for j := b.topLeft.Col; i <= b.bottomRight.Col; j++ {
			if inner, ok := b.Alive[i]; ok {
				if _, ok := inner[j]; ok {
					fmt.Print("*")
					continue
				}
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func initBoard(seed []Cell) Board {
	bottomRight := seed[0]
	topLeft := seed[0]
	init := make(map[int]map[int]interface{})
	for _, c := range seed {
		init[c.Row][c.Col] = nil
		if c.Row > bottomRight.Row {
			bottomRight.Row = c.Row
		}
		if c.Col > bottomRight.Col {
			bottomRight.Col = c.Col
		}
		if c.Row < topLeft.Row {
			topLeft.Row = c.Row
		}
		if c.Col < topLeft.Col {
			topLeft.Col = c.Col
		}
	}
	return Board{init, topLeft, bottomRight}
}

func nextTurn(prev Board) Board {

}
