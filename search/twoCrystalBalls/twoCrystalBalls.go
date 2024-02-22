package main

import (
	"fmt"
	"math"
)

func two_crystal_balls(breaks []bool) int {
	interval := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := interval

	for i < len(breaks) {
		if (breaks[i]) == true {
			break
		}
		i += interval
	}

	i -= interval
  
  j := 0

	for j < interval && i < len(breaks) {
		if breaks[i] {
			return i
		}
    i++
    j++
	}

	return -1
}

func main() {
	breaks := []bool{false, false, false, false, true, true, true, true, true}
	fmt.Println(two_crystal_balls(breaks))
}
