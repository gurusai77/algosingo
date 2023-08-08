package main

import (
	"fmt"
	"math/bits"
)

func secondl(arr []int) int {
	fl, sl := arr[0], -1
	for _, v := range arr {
		if v > fl {
			fl, sl = v, fl
		} else if v < fl && v > sl {
			sl = v
		}
	}
	return sl
}

func secondsmallest(arr []int) int {
	fs, ss := arr[0], (1<<bits.UintSize)/2-1
	for _, v := range arr {
		if v < fs {
			fs, ss = v, fs
		} else if v != fs && v < ss {
			ss = v
		}
	}
	return ss
}

func removeDuplicates(arr []int) ([]int, int) {
	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			arr[i+1] = arr[j]
			i++
		}
	}
	return arr, i
}

/*
this solution handles negative numbers also
*/
func main() {
	arr := []int{100, 5, 2, 3, 3, 6, 4, 1, 1, -1, -10}
	fl := secondl(arr)
	ss := secondsmallest(arr)
	fmt.Println(fl)
	fmt.Println(ss)
	arr1 := []int{-10, 1, 1, -2, -10}
	fmt.Println(secondl(arr1))
	arr2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	arr2, i := removeDuplicates(arr2)
	fmt.Println(arr2, i)
	// make other numbers 0
	for j := i; j < len(arr2)-1; j++ {
		arr2[j+1] = 0
	}
	fmt.Println(arr2)
}
