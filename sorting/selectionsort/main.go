package main

import "fmt"

//Advantages of Selection Sort Algorithm
//Simple and easy to understand.
//Works well with small datasets.

//Disadvantages of the Selection Sort Algorithm
//Selection sort has a time complexity of O(n^2) in the worst and average case.
//Does not work well on large datasets.
//Does not preserve the relative order of items with equal keys which means it is not stable.

//find and selects the index of min element and swap with current element in outer loop

// worst and avg case o(n2)
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
