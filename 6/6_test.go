package euler

import "testing"

func TestProblem(t *testing.T) {
	result := sumSquareDifference()
	if result != 25164150 {
		t.Error("Incorrect result, expected 232792560 got ", result)
	}
}
