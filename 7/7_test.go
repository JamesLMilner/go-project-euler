package euler

import "testing"

func TestProblem(t *testing.T) {
	result := nthPrimeNumber(10001)
	if result != 104743 {
		t.Error("Incorrect result, expected 104743 got ", result)
	}
}
