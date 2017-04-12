package euler

import "testing"

func TestProblem(t *testing.T) {
	result := smallestMultiple()
	if result != 232792560 {
		t.Error("Incorrect result, expected 232792560 got ", result)
	}
}
