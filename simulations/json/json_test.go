package json

import "testing"

func TestJSON_Run(t *testing.T) {
	expected := `{"X":0,"Y":1}`
	s, err := New([]byte(`{"Width":4,"Height":4, "X":2, "Y":2, "Commands": ["f", "l", "f", "r", "b", "r", "b", "l", "f", "q"]}`))
	if err != nil {
		t.Error(err)
	}
	result, err := s.Run()
	if err != nil {
		t.Error(err)
	}
	if string(result) != expected {
		t.Errorf("Wrong solution: %s instead of %s", string(result), expected)
	}
}
