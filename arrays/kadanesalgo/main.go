package main

import (
	"fmt"
	"math"
)

func kadane(arr []int) (int, int, int) {
	maxSoFar := math.MinInt
	sum := 0
	start := 0
	end := 0
	s := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > maxSoFar {
			maxSoFar = sum
			start = s
			end = i
		}
		if sum < 0 {
			sum = 0
			s = i + 1
		}
	}
	return maxSoFar, start, end
}

// kadanes algo
func main() {
	arr := []int{-1, -2, 5, -9, 2, 3, 4, 7, -6}
	fmt.Println(kadane(arr))
}
