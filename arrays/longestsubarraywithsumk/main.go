package main

import (
	"fmt"
	"math"
)

func longestarray(a []int, k int) int {
	sum := a[0]
	maxLen := 0
	j := 0
	for i := 0; i < len(a); {
		for sum > k {
			sum -= a[j]
			j++
		}
		if sum == k {
			maxLen = int(math.Max(float64(maxLen), float64(i-j+1)))
		}
		i++
		if i < len(a) {
			sum += a[i]
		}
	}
	return maxLen
}

func longestarray2(a []int, k int) int {
	sum := 0
	maxLen := 0
	j := 0
	for i := 0; i < len(a); {
		sum += a[i]
		if sum < k {
			i++
		} else if sum == k {
			maxLen = int(math.Max(float64(maxLen), float64(i-j+1)))
			i++
		} else if sum > k {
			for j < len(a) && sum > k {
				sum -= a[j]
				j++
			}
			if sum == k {
				maxLen = int(math.Max(float64(maxLen), float64(i-j+1)))
			}
			i++
		}

	}
	return maxLen
}

// 2 pointer approach
func main() {
	fmt.Println(longestarray([]int{1, 2, 3, 4, 5, 6}, 9))
	fmt.Println(longestarray2([]int{1, 2, 3, 4, 5, 6}, 9))
}
