package euler
import "testing";

func TestProblem(t *testing.T) {
   if (findSum() != 233168) {
     t.Error("Incorrect result, expected 233168");
   }
}
