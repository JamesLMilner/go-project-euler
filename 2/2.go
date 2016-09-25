package euler
import "fmt";

func main() {
  fmt.Println(findSum());
}

func findSum() int {


   f := fibonacci();
   sum := 0;

   for true {
      v := f();
      if (v < 4000000) {
        if (v % 2 == 0) {
          sum += v;
        }
      } else {
        break;
      }
   }

   return sum;
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x + y
		return x
	}
}
