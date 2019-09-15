package board

import "testing"

func TestNew(t *testing.T) {
	m, err := New(4, 4)
	if err != nil {
		t.Error(err)
	}
	if m.Height != 4 || m.Width != 4 {
		t.Error("Failed to set width and height correctly")
	}
	for r, column := range m.Cells {
		for c, cell := range column {
			if cell != true {
				t.Errorf("Failed to initialise cell [%d,%d] correctly", c, r)
			}
		}
	}

	_, err = New(-1, 2)
	if err == nil {
		t.Errorf("Failed to detect wrong board size")
	}
}

func TestBoard_IsOnTheBoard(t *testing.T) {
	m, err := New(4, 4)
	if err != nil {
		t.Error(err)
	}
	if !m.IsOnTheBoard(1, 1) {
		t.Errorf("Wrongly reported cell [1,1] as offboard")
	}
	if m.IsOnTheBoard(8, 8) != false {
		t.Errorf("Failed to detect cell [8,8] as offboard")
	}
}
