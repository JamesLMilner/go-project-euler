package euler

import (
  "fmt"
  "strconv"
)

func main() {
  fmt.Println(getLargestPalindromicNumber());
}

func getLargestPalindromicNumber() int {

  largest := 0

  for i := 100; i < 1000; i++ {
      for j := 100; j < 1000; j++ {
        num := i * j
        rev := strconv.FormatInt(int64(num), 10)
        if (Reverse(rev) == rev && num > largest) {
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
