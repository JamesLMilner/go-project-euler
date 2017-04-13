package euler

import "testing"

func TestProblem(t *testing.T) {
	result := pythagoreanTriplet()
	if result != 31875000 {
		t.Error("Incorrect result, expected 31875000 got ", result)
	}
}
