package euler

import "testing"

func TestProblem(t *testing.T) {
	result := largestProduct()
	if result != 23514624000 {
		t.Error("Incorrect result, expected 23514624000 got ", result)
	}
}
