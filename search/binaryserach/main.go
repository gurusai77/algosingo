package main

import "fmt"

// iterative
func bs(a []int, target int) int {
	low := 0
	high := len(a)
	for low < high {
		mid := (low + high) / 2
		if target < a[mid] {
			high = mid - 1
		} else if target > a[mid] {
			low = mid + 1
		} else if target == a[mid] {
			return mid
		}
	}
	return -1
}

// recursive
func bsr(a []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if target == a[mid] {
		return mid
	} else if target < a[mid] {
		return bsr(a, low, mid-1, target)
	} else if target > a[mid] {
		return bsr(a, mid+1, high, target)
	}
	return -1
}

func main() {
	fmt.Println(bs([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2))
	fmt.Println(bsr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 9, 3))
}
