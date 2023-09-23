package main

import "fmt"

// two solutions as below

//items are not sorted
//Space complexity must be O(1): brute force
//Space complexity can be O(n): hashmap/set

//Items are sorted: two pointers

// two pointer
func adddigitstotarget(a []int, target int) (x [][]int) {
	// if array is not sorted this solution fails
	if len(a) <= 1 {
		return [][]int{}
	}

	ptr1, ptr2 := 0, len(a)-1
	for ptr1 < ptr2 {
		sum := a[ptr2] + a[ptr1]
		if target == sum {
			x = append(x, []int{ptr1, ptr2})
		}
		if sum > target {
			ptr2--
		} else {
			ptr1++
		}
	}
	return x
}

func twoSum(nums []int, target int) (x [][]int) {
	// Space: O(n)
	s := make(map[int]int)

	// Time: O(n)
	for idx, num := range nums {
		// Time: O(1)
		if pos, ok := s[target-num]; ok {
			x = append(x, []int{pos, idx})
		}
		s[num] = idx
	}
	return x
}

func main() {
	fmt.Println(adddigitstotarget([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 13))
}
