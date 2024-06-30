package main

import "fmt"

//worst and avg case o(n2)
// sort until the elements is rightly placed in correct position

func insert(a []int) []int {
	for i := 1; i < len(a); i++ {
		temp := a[i]
		j := i - 1
		fmt.Printf("temp values:%v\n", temp)
		fmt.Printf("values of j:%v\n", j)
		for j >= 0 && a[j] > temp {
			a[j+1] = a[j]
			j = j - 1
			fmt.Printf("intermediate updated array:%v\n", a)
		}
		a[j+1] = temp
		fmt.Printf("updated array:%v\n", a)
	}
	return a
}

func main() {
	fmt.Println(insert([]int{23, 6, 8, 34, 22, 19, 12}))
}
