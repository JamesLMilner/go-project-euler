package euler

import (
	"fmt"
)

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

func main() {
	fmt.Println(summationOfPrimes(2000000))
}

func summationOfPrimes(n int) int {

	sum := 0
	primes := sieveOfEratosthenes(n)
	for i := 2; i < len(primes); i++ {
		if primes[i] {
			sum += i
		}
	}
	return sum

}

func sieveOfEratosthenes(n int) []bool {

	//  Input: an integer n > 1.

	//  Let A be an array of Boolean values, indexed by integers 2 to n,
	//  initially all set to true.

	//  for i = 2, 3, 4, ..., not exceeding √n:
	//    if A[i] is true:
	//      for j = i2, i2+i, i2+2i, i2+3i, ..., not exceeding n:
	//        A[j] := false.

	//  Output: all i such that A[i] is true.

	sieve := []bool{}
	for i := 0; i < n; i++ {
		sieve = append(sieve, true)
	}

	for i := 2; i < n; i++ {
		if sieve[i] {
			for j, k := i*i, 1; j < n; j, k = i*i+k*i, k+1 {
				sieve[j] = false
			}
		}
	}
	return sieve
}
