package main

import (
	"fmt"
	"math"
)

func sequence(a []int) float64 {
	var count float64 = 1
	var maxCount float64 = -1
	for i := 0; i < len(a)-1; i++ {
		if a[i+1]-a[i] == 1 {
			count++
		} else if a[i+1]-a[i] == 0 {
			continue
		} else {
			count = 1
		}
		maxCount = math.Max(maxCount, count)
	}
	return maxCount
}

func main() {
	fmt.Println(sequence([]int{1, 1, 2, 2, 3, 3, 3, 5, 4, 4, 4, 6, 5, 9}))
}
