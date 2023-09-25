package main

import "fmt"

func merge(a []int, b []int) []int {
	var c = make([]int, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c[k] = a[i]
			i++
			k++
		} else {
			c[k] = b[j]
			j++
			k++
		}
	}
	for i < len(a) {
		c[k] = a[i]
		i++
		k++
	}

	for j < len(b) {
		c[k] = b[j]
		j++
		k++
	}

	return c
}

func main() {
	fmt.Println(merge([]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 10}))
}
