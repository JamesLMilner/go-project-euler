package euler

import (
  "fmt"
  "math"
  "sort"
)

func main() {
  fmt.Println(getLargestPrimeFactor());
}

func getLargestPrimeFactor() int {
   num := 600851475143
   factors := primeFactors(num);
   sortedFactors := sort.IntSlice(factors)
	 sort.Sort(sortedFactors)

   return sortedFactors[len(sortedFactors)-1]
}

func primeFactors(num int) []int {

  var factors []int
  f := float64(num)

  for i := 2; i < int(math.Sqrt(f)) + 1; i++ {

     if (num % i == 0) {

       fact := num / i

       if (isPrime(fact)) {
        factors = append(factors, i)
       }
       if (isPrime(i)) {
        factors = append(factors, i)
       }

     }

  }

  return factors

}

func isPrime(num int) bool {

  prime := true;
  f := float64(num)

  for i := 2; i < int(math.Sqrt(f)) + 1; i++ {
     if (num % i == 0) {
       prime = false;
     }
  }

  return prime

}
