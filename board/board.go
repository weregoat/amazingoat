package board

import (
	"errors"
	"fmt"
)

/*
The board is *always* a matrix of rows and columns.
Each row has the same number of columns.
To allow for different shapes the single cells may in the board (true) or not (false), like holes.
Example:
|  *****  |
|*  ***  *|
|**  *  **|
|***   ***|
Thus, making a "V" maze (* = holes)
I think easier this way than to try some other way. Didn't think too hard about it but maybe (Double linked blocks,
I suppose, which Golang doesn't have and I would have to implement myself :).

Remember (useful to understand why the Unit Circle):
Moving North and South changes row; -1 and +1 respectively.
Moving East and West changes column; respectively +1 and -1.
*/

// Maze is the columnsXrows board the maze created on.
type Board struct {
	Width  int      // Number of columns
	Height int      // Number of cells
	Cells  [][]bool // True if the cells belong to the board
}

// New creates a new board width*height
func New(width, height int) (Board, error) {
	var err error
	m := Board{Width: width, Height: height}
	if width < 0 || height < 0 { // Didn't bother to enforce an upper limit
		err = errors.New(fmt.Sprintf("invalid board width (%d) or height (%d)", width, height))
	} else {
		cells := make([][]bool, width)
		for i := range cells {
			cells[i] = make([]bool, height)
			for j := 0; j < height; j++ {
				cells[i][j] = true
			}

		}
		m.Cells = cells
	}
	return m, err
}

// IsOnTheBoard checks if the cell at the given coordinates are inside the boundaries and "in" the board (value true).
func (b *Board) IsOnTheBoard(x, y int) bool {
	in := false
	if x >= 0 && x < b.Width && y >= 0 && y < b.Height {
		in = b.Cells[x][y]
	}
	return in
}

/* These are just here to point out a way to handle making or fixing "holes" in the board. But it would require
to implement a more complex protocol, where you have to define how to communicate the holes (an ascii file to read, an
array of coordinates in the JSON input...)
// Add a cell to the board.
func (b *Board) Add(x,y int) {
	if x >= 0 && x < b.Width && y >= 0 && y < b.Height {
		b.Cells[x][y] = true
	}
}

// Remove a cell from the board.
func (b *Board) Remove(x,y int) {
	if x >= 0 && x < b.Width && y >= 0 && y < b.Height {
		b.Cells[x][y] = false
	}
}
*/
