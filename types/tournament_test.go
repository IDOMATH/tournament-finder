package types

import "testing"

func AgeDivisionArrayToIntTest(t *testing.T) {
	expected := 0
	tourney := Tournament{AgeDivision: [8]bool{0, 0, 0, 0, 0, 0, 0, 0}}

	got := tourney.AgeDivisionArrayToInt()

	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}
