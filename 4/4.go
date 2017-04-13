package euler

import (
	"fmt"
	"strconv"
)

// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

func main() {
	fmt.Println(getLargestPalindromicNumber())
}

func getLargestPalindromicNumber() int {

	largest := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			num := i * j
			rev := strconv.FormatInt(int64(num), 10)
			if Reverse(rev) == rev && num > largest {
				largest = num
			}
		}
	}

	return largest

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
