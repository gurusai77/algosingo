package main

import "fmt"

//Advantages of Bubble Sort:
//Bubble sort is easy to understand and implement.
//It does not require any additional memory space.
//It is a stable sorting algorithm, meaning that elements with the same key value maintain their relative order in the sorted output.
//Disadvantages of Bubble Sort:
//Bubble sort has a time complexity of O(N2) which makes it very slow for large data sets.
//Bubble sort is a comparison-based sorting algorithm, which means that it requires a comparison operator to determine the relative order of elements in the input data set. It can limit the efficiency of the algorithm in certain cases.

// best case is o(n), worst case is o(n2)

func bubble(a []int) []int {
	for i := len(a) - 1; i >= 0; i-- {
		swap := false
		for j := 0; j <= i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
	return a
}

func main() {
	fmt.Println(bubble([]int{85, 9, 2, 34, 76, 67, 45}))
}
