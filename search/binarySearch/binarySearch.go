package main

import (
  "fmt"
  "math"
)

func binary_search(haystack []int, needle int) bool {
  
  lo := 0
  hi := len(haystack)

  for lo < hi {
    m := int(math.Floor(float64(lo + (hi - lo) / 2)))

    v := haystack[m]

    if (v == needle) {
      return true
    } else if (needle < v) {
      hi = m
    } else {
      lo = m + 1
    }
  }

  return false
}

func main () {
  haystack := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
  needle := 7

  fmt.Println(binary_search(haystack, needle))
}
