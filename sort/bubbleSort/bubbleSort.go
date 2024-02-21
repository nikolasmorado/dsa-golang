package main

func bubble_sort(arr []int) []int {
  for i := 0; i < len(arr); i++ {
    for j := 0; j < len(arr) - 1 - i; j++ {
      if arr[j] > arr[j + 1] {
        arr[j], arr[j + 1] = arr[j + 1], arr[j]
      }
    }
  }
  return arr
}

func main() {
  arr := []int{3, 2, 1, 5, 4}
  arr = bubble_sort(arr)
  for i := 0; i < len(arr); i++ {
    print(arr[i], " ")
  }
  println()
}

