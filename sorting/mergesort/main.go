package main

import "fmt"

//O(N log(N)),  Merge Sort is a recursive algorithm and time complexity can be expressed as following recurrence relation.
//Auxiliary Space: O(N), In merge sort all elements are copied into an auxiliary array. So N auxiliary space is required for merge sort.

func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func divide(a []int) []int {
	if len(a) < 2 {
		return a
	}
	first := divide(a[:len(a)/2])
	second := divide(a[len(a)/2:])
	return merge(first, second)
}

func main() {
	a := []int{2, 1, 4, 3, 21, 19, 13, 16, 11}
	fmt.Println(divide(a))
}
