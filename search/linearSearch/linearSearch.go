package main

import "fmt"

func linear_search(haystack []int, needle int) bool {
  for i := 0; i < len(haystack); i++ {
    if haystack[i] == needle {
      return true
    }
  }
  return false
}

func main() {
  arr := []int{1, 2, 5, 6, 8}

    res := linear_search(arr, 5)

    fmt.Println(res)

}
