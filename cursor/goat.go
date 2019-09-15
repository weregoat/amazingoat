package cursor

import (
	"errors"
	"fmt"
	"github.com/weregoat/amazingoat/board"
	"math"
)

// 90° facing for string
const N = "north"

// 270° facing for string
const S = "south"

// 0° facing for string
const E = "east"

// 180° facing for string
const W = "west"

// The rotation step as radian.
const RSTEP = math.Pi / 2 // 90° or π½
// The initial orientation as radians.
const RSTART = RSTEP // A coincidence of sort.

// Goat represent a horned cursor exploring the maze.
type Goat struct {
	// Properties are public as I try to follow the Python idea that we are all consenting adults.
	Board  board.Board // The maze the goat is in.
	X      int         // Column
	Y      int         // Row
	Rad    float64     // Facing angle (it should be a multiple of RSTEP)
	Cosine int         // The cosine of the angle (define direction along the X axis)
	Sine   int         // The sine of the angle (define direction along the Y axis)
	Facing string      // Facing as a string "North" or "N" etc.

}

// New creates a new cursor (as a goat) at the given coordinates and facing North (0,1)
func New(maze board.Board, x, y int) (Goat, error) {
	var err error
	goat := Goat{Board: maze, X: x, Y: y}
	if !maze.IsOnTheBoard(x, y) {
		err = errors.New(fmt.Sprintf("the starting position [%d,%d] is not on the board", x, y))
	} else {
		err = goat.Orientate(RSTART) // Face it North
	}
	return goat, err
}

// String prints the position of the cursor as a string.
func (goat *Goat) String() string {
	return fmt.Sprintf("goat at [%d,%d] facing %s", goat.X, goat.Y, goat.Facing)
}

// Forward moves the cursor forward according to the orientation.
func (goat *Goat) Forward() {
	x := goat.X + goat.Cosine // It moves in the direction is facing so previous column (-1) if W (-1) and next (+1) if E (1)
	y := goat.Y - goat.Sine   // Previous row (-1) if N (1) and next row (+1) if S (-1)
	// The goat steps out of the board
	if !goat.Board.IsOnTheBoard(x, y) {
		goat.X = -1
		goat.Y = -1
		return
	}
	goat.X = x
	goat.Y = y
}

// Back moves the cursor backwards according to the orientation.
func (goat *Goat) Back() {
	x := goat.X - goat.Cosine // It moves in the opposite direction it's facing so next column (+1) if W (-1) and previous (-1) if E (1)
	y := goat.Y + goat.Sine   // It moves in the opposite direction it's facing so next row (+1) if N (1) and previous (-1) if S (-1)
	if !goat.Board.IsOnTheBoard(x, y) {
		goat.X = -1
		goat.Y = -1
		return
	}
	goat.X = x
	goat.Y = y
	return
}

// Right rotates the angle clockwise by subtracting 90 degrees to the orientation.
func (goat *Goat) Right() error {
	angle := goat.Rad - RSTEP
	return goat.Orientate(angle)
}

// Left rotates the angle counterclockwise by adding 90 degrees to the orientation.
func (goat *Goat) Left() error {
	angle := goat.Rad + RSTEP
	return goat.Orientate(angle)
}

// Orientate sets the sine and cosine according to the given radians
func (goat *Goat) Orientate(rad float64) error {
	var err error
	cosine := int(math.Cos(rad)) // Given we the angle is changed by PI/2 the result is always -1,0, or 1
	sine := int(math.Sin(rad))
	facing := facing(cosine, sine)
	if len(facing) == 0 {
		err = errors.New(fmt.Sprintf("invalid angle: %f.2º", (math.Pi/180)*rad))
	} else {
		goat.Cosine = cosine
		goat.Sine = sine
		goat.Facing = facing
		goat.Rad = rad
	}
	return err
}

func facing(cosine, sine int) string {
	// I am using the Unit circle to orient the cursor.
	// See https://en.wikipedia.org/wiki/Unit_circle#/media/File:Unit_circle_angles_color.svg
	// This way I can use the sine and cosine value to pick the next coordinates
	// and I don't need to reset the angle (it doesn't matter which multiple of Pi/2 we are at)
	var facing string
	switch {
	case cosine == 0 && sine == 1: // (0,1) 90º
		facing = N
	case cosine == 1 && sine == 0: // (1,0) 0º
		facing = E
	case cosine == 0 && sine == -1: // (0,-1) 270º
		facing = S
	case cosine == -1 && sine == 0: // (-1,0) 180º
		facing = W
	}
	// If the sine and cosine are wrong, the facing is the empty string.
	return facing
}
