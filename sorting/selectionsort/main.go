package main

import "fmt"

func selection(a []int) []int {
	for i := 0; i < len(a); i++ {
		// store the ith index as min element
		minIdx := i
		// find minimum element
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
			// sort with the ith element
			a[i], a[minIdx] = a[minIdx], a[i]
		}
	}
	return a
}

func main() {
	fmt.Println(selection([]int{56, 44, 67, 78, 23, 9}))
}
