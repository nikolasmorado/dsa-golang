package main

import "fmt"

func partition(arr []int, lo, hi int) int {
    pivot := arr[hi]
    idx := lo - 1

    for i := lo; i < hi; i++ {
        if arr[i] <= pivot {
            idx++
            arr[i], arr[idx] = arr[idx], arr[i]
        }
    }

    idx++
    arr[hi], arr[idx] = arr[idx], arr[hi]

    return idx
}

func quick_sort(arr []int, lo, hi int) {
    if lo >= hi {
        return
    }

    pivotIdx := partition(arr, lo, hi)

    quick_sort(arr, lo, pivotIdx-1)
    quick_sort(arr, pivotIdx+1, hi)
}

func main() {
    arr := []int{4, 5, 2, 8, 1, 2, 10, 32, 6, 100, 22, 400}

    quick_sort(arr, 0, len(arr)-1)

    for _, value := range arr {
        fmt.Print(value, " ")
    }
    fmt.Println()
}

