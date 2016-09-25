package euler
import "testing";

func TestProblem(t *testing.T) {
  result := getLargestPalindromicNumber()
   if (result  != 906609) {
     t.Error("Incorrect result, expected 906609 got ", result);
   }
}
