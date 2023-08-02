package main

import (
	"fmt"
	"math"
)

func getLeaders(arr []int) []int {
	max := math.MinInt
	i := 0
	var out []int
	for i = len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			max = arr[i]
			out = append(out, arr[i])
		}
	}
	return out
}

func main() {
	arr := []int{20, 3, 4, 5, 1, 0, -1}
	fmt.Println(getLeaders(arr))
}
