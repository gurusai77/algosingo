package main

import "fmt"

func removeDuplicates(nums []int) int {
	var s int
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[s] {
			continue
		}
		s++
		nums[s] = nums[i]
	}
	return s + 1
}
func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 3, 3}))
}
