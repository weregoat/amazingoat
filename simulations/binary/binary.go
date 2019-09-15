package binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/weregoat/amazingoat/board"
	"github.com/weregoat/amazingoat/cursor"
	"log"
)

// Constant defining the command value for Quit.
const QUIT int = 0

// Constant defining the command value for forward.
const FORWARD int = 1

// Constant defining the command value for backward.
const BACKWARD int = 2

// Constant defining the command value for rotating clockwise.
const RIGHT int = 3

// Constant defining the command value for rotating counterclockwise.
const LEFT int = 4

// Binary define the struct for a simulation processing in and out as binary streams.
type Binary struct {
	Board    board.Board
	Goat     cursor.Goat
	Commands []int
	debug    bool
}

// New creates a new simulation from the given stream according to its specific rules.
// "
// ● The size of the table as two 16-bit little-endian integers [width, height]
// ● The objects starting position as two 16-bit little-endian integers [x, y]
//
// This is followed by an arbitrarily long stream of commands of 8-bit integers."
//
func New(payload []byte) (s Binary, err error) {
	var commands []int
	// 2 bytes for the Width
	// 2 bytes for the Height
	// 2 bytes for the starting X
	// 2 bytes for the starting Y
	// Makes 8 at least bytes
	if len(payload) < 8 {
		err = errors.New("not enough bytes to initialise the simulation")
		return
	}
	width := int(binary.LittleEndian.Uint16(payload[:2]))
	height := int(binary.LittleEndian.Uint16(payload[2:4]))
	x := int(binary.LittleEndian.Uint16(payload[4:6]))
	y := int(binary.LittleEndian.Uint16(payload[6:8]))
	for _, b := range payload[8:] {
		commands = append(commands, int(b))
	}
	maze, err := board.New(width, height)
	if err != nil {
		return
	}
	goat, err := cursor.New(maze, x, y)
	if err != nil {
		return
	}
	s.Goat = goat
	s.Board = maze

	s.Commands = commands
	return
}

// Run the simulation
func (s *Binary) Run() ([]byte, error) {
	var err error
	s.log(s.Goat.String())
	s.log(fmt.Sprintf("commands to process: %v", s.Commands))
commandsExecution: // https://golang.org/ref/spec#Break_statements
	for _, command := range s.Commands {
		s.log(fmt.Sprintf("processing command %d", command))
		switch command {
		case FORWARD:
			s.Goat.Forward()
		case BACKWARD:
			s.Goat.Back()
		case RIGHT:
			err = s.Goat.Right()
		case LEFT:
			err = s.Goat.Left()
		case QUIT:
			break commandsExecution
		default:
			err = errors.New(fmt.Sprintf("unknown command: %d", command))
		}
		if err != nil {
			s.log(err.Error())
		}
		s.log(s.Goat.String())
		if s.Goat.X == -1 && s.Goat.Y == -1 {
			break
		}
	}
	var result []byte
	bx := make([]byte, 2)
	by := make([]byte, 2)
	binary.LittleEndian.PutUint16(bx, uint16(s.Goat.X))
	binary.LittleEndian.PutUint16(by, uint16(s.Goat.Y))
	result = []byte{bx[0], bx[1], by[0], by[1]}
	return result, err
}

// Debug sets the debug property for logging the cursor movements
func (s *Binary) Debug(debug bool) {
	s.debug = debug
}

// log sends the given message to logging only if debug is enabled.
func (s *Binary) log(message string) {
	if s.debug {
		log.Print(message)
	}
}
