package main

import (
	"fmt"
	"math"
)

// this is the best solution with o(n) time complexity and O(n) space complexity
// if dont use array the space complexity is O(1)

// consider max as minInt and start from right side of array, so if current number is greater than max then update max
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
	arr := []int{20, 3, 4, 5, 1, 0, -1, 7}
	fmt.Println(getLeaders(arr))
}
