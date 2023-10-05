package main

import "fmt"

func removeDuplicates(nums []int) int {
	var unique int
	for i := 2; i < len(nums); i++ {
		// if nums[i] == nums[i-2] && nums[i] == nums[i-1] {
		if nums[i] == nums[i-2] {
			nums[i] = 0
			continue
		}
		unique++
	}
	fmt.Println(nums)
	return unique
}
func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 2, 2, 3, 3, 3, 3}))
}
