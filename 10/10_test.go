package euler

import "testing"

func TestProblem(t *testing.T) {
	result := summationOfPrimes(2000000)
	if result != 142913828922 {
		t.Error("Incorrect result, expected 142913828922 got ", result)
	}
}
