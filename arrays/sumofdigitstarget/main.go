package main

import "fmt"

//tems are not sorted
//Space complexity must be O(1): brute force
//Space complexuty can be O(n): hashmap/set
//Items are sorted: two pointers

func adddigitstotarget(a []int, target int) (int, int) {
	// if array is not listed this solution fails

	if len(a) <= 1 {
		return 0, 0
	}

	ptr1, ptr2 := 0, len(a)-1
	for ptr1 < ptr2 {
		sum := a[ptr2] + a[ptr1]
		if target == sum {
			return ptr1, ptr2
		} else if sum > target {
			ptr2--
		} else if sum < target {
			ptr1++
		}
	}
	return 0, 0
}

func twoSum(nums []int, target int) []int {
	// Space: O(n)
	s := make(map[int]int)

	// Time: O(n)
	for idx, num := range nums {
		// Time: O(1)
		if pos, ok := s[target-num]; ok {
			return []int{pos, idx}
		}
		s[num] = idx
	}
	return []int{}
}

func main() {
	fmt.Println(adddigitstotarget([]int{1, 4, 5, 8, 5, 4, 6}, 9))
	fmt.Println(twoSum([]int{1, 4, 5, 8, 5, 4, 6}, 13))
}
