package euler

import (
	"fmt"
)

// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

func main() {
	fmt.Println(smallestMultiple())
}

func smallestMultiple() int {

	searching := true
	i := 0
	var result int

	for searching {

		i++

		for j := 1; j <= 20; j++ {

			if i%j == 0 {
				if j == 20 {
					searching = false
					result = i
				}
			} else {
				break
			}

		}

	}

	return result
}
