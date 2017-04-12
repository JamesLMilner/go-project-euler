package euler

import (
	"fmt"
)

// The sum of the squares of the first ten natural numbers is,

// 1**2 + 2**2 + ... + 10**2 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)**2 = 55**2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func main() {
	fmt.Println(nthPrimeNumber(10001))
}

// Super niave nth prime number function
func nthPrimeNumber(n int) int {

	num := 3
	primeCount := 0
	currentPrime := 2
	var isPrime bool

	for primeCount != n {

		isPrime = true

		// Using Square route would be more efficent
		for i := 2; i < num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primeCount++
			currentPrime = num
		}

		num++

	}

	return currentPrime

}
