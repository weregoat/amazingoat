package cursor

import (
	"github.com/weregoat/amazingoat/board"
	"math"
	"testing"
)

const X1 = 2
const Y1 = 2
const FACING1 = "north"

func TestGoat_New(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	if goat.X != X1 {
		t.Errorf("Wrong starting X: %d", goat.X)
	}

	if goat.Y != Y1 {
		t.Errorf("Wrong starting Y: %d", goat.X)
	}

	if goat.Facing != FACING1 {
		t.Errorf("Wrong initial facing: %s", goat.Facing)
	}
}

func TestGoat_Back(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	goat.Back()
	if goat.Y != (Y1+1) || goat.X != X1 {
		t.Errorf("Failed to move backwards correctly")
	}
}

func TestGoat_Forward(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	goat.Forward()
	if goat.Y != (Y1-1) || goat.X != X1 {
		t.Errorf("Failed to move forwards correctly")
	}
}

func TestGoat_Left(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	goat.Left()
	goat.Back()
	if goat.Y != Y1 || goat.X != (X1+1) {
		t.Errorf("Failed to rotate left correctly")
	}
}

func TestGoat_Right(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	goat.Right()
	goat.Forward()
	if goat.Y != Y1 || goat.X != (X1+1) {
		t.Errorf("Failed to rotate right correctly")
	}
}

func TestGoat_FallingOff(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	goat.Forward()
	goat.Forward()
	goat.Forward()
	if goat.Y != -1 || goat.X != -1 {
		t.Errorf("Failed to detect off board position")
	}
}

func TestGoat_Orientate(t *testing.T) {
	maze, err := board.New(4, 4)
	if err != nil {
		t.Error(err)
	}
	goat, err := New(maze, X1, Y1)
	if err != nil {
		t.Error(err)
	}
	err = goat.Orientate(math.Pi) // 180°
	if err != nil {
		t.Error(err)
	}
	if goat.Rad != math.Pi || goat.Cosine != -1 || goat.Sine != 0 || goat.Facing != "west" {
		t.Errorf("Failed to orientate correctly for 180°")
	}
	err = goat.Orientate(math.Pi / 3) // Invalid
	if err == nil {
		t.Errorf("Failed to exit on invalid angle")
	}
}
