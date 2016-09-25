package euler
import "testing";

func TestProblem(t *testing.T) {
  result := getLargestPrimeFactor()
   if (result  != 6857) {
     t.Error("Incorrect result, expected 6857 got ", result);
   }
}
