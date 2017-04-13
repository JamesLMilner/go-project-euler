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
	fmt.Println(sumSquareDifference())
}

func sumSquareDifference() int {

	sqsum := 0
	sum := 0

	for j := 1; j <= 100; j++ {

		sqsum += j * j
		sum += j

	}

	sum = sum * sum

	return sum - sqsum

}
