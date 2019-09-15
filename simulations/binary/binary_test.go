package binary

import "testing"

func TestJSON_Run(t *testing.T) {
	expected := []byte{0x0, 0x0, 0x1, 0x0} // 0,1
	s, err := New([]byte{0x4, 0x0, 0x4, 0x0, 0x2, 0x0, 0x2, 0x0, 0x1, 0x4, 0x1, 0x3, 0x2, 0x3, 0x2, 0x4, 0x1, 0x0})
	if err != nil {
		t.Error(err)
	}
	result, err := s.Run()
	if err != nil {
		t.Error(err)
	}
	if string(result) != string(expected) {
		t.Errorf("Wrong solution: %v instead of %v", result, expected)
	}
}
