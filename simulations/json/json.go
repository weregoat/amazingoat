package json

/*
	Alternative simulator to show how the binary can be altered to handle different specifications.
    In this case:
    "Change the binary form of the protocol to JSON"
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/weregoat/amazingoat/board"
	"github.com/weregoat/amazingoat/cursor"
	"log"
)

const QUIT string = "q"
const FORWARD string = "f"
const BACKWARD string = "b"
const RIGHT string = "r"
const LEFT string = "l"

type JSON struct {
	Maze     board.Board
	Goat     cursor.Goat
	Commands []string
	debug    bool
}

type Result struct {
	X int
	Y int
}

// {"Width:4","Height:4", "X:2", "Y:2", "Commands": ["f", "l", "f", "", "b", "q"]}
type body struct {
	Width    int
	Height   int
	X        int
	Y        int
	Commands []string
}

func New(payload []byte) (s JSON, err error) {
	b := body{}
	err = json.Unmarshal(payload, &b)
	if err != nil {
		return
	}
	maze, err := board.New(b.Width, b.Height)
	if err != nil {
		return
	}
	goat, err := cursor.New(maze, b.X, b.Y)
	if err != nil {
		return
	}
	s.Goat = goat
	s.Maze = maze
	s.Commands = b.Commands
	return
}

func (s *JSON) Run() ([]byte, error) {
	var err error
	s.log(s.Goat.String())
	s.log(fmt.Sprintf("commands to process: %v", s.Commands))
commandsExecution: // https://golang.org/ref/spec#Break_statements
	for _, command := range s.Commands {
		s.log(fmt.Sprintf("processing command %s", command))
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
			err = errors.New(fmt.Sprintf("unknown command: %s", command))
		}
		if err != nil {
			s.log(err.Error())
		}
		s.log(s.Goat.String())
		if s.Goat.X == -1 && s.Goat.Y == -1 {
			break
		}
	}

	position := Result{s.Goat.X, s.Goat.Y}
	result, jsonError := json.Marshal(position)
	if jsonError != nil {
		return result, jsonError
	}
	return result, err
}

func (s *JSON) Debug(debug bool) {
	s.debug = debug
}

func (s *JSON) log(message string) {
	if s.debug {
		log.Print(message)
	}
}
