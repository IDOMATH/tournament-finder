package types

import "testing"

func AgeDivisionArrayToIntTest(t *testing.T) {
	expected := 0
	tourney := Tournament{AgeDivision: [8]bool{false, false, false, false, false, false, false, false}}

	got := tourney.AgeDivisionArrayToInt()

	if got != expected {
		t.Errorf("Expected %v, got: %v", expected, got)
	}
}
