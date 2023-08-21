package main

import "fmt"

//worst and avg case o(n2)
// sort until the elements is rightly placed in correct position

func insert(a []int) []int {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i - 1
		for j >= 0 && a[j] > temp {
			a[j+1] = a[j]
			j = j - 1
		}
		a[j+1] = temp
	}
	return a
}

func main() {
	fmt.Println(insert([]int{23, 6, 8, 34, 22, 19, 12}))
}
