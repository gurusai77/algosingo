package main

import "fmt"

// import (
//
//	"fmt"
//	"math"
//
// )
//
// O(n2) we can optimize better this approach
func majoritybrute(a []int) (int, int) {
	for i := 0; i < len(a); i++ {
		var count int = 0
		for j := 1; j < len(a); j++ {
			if a[j] == a[i] {
				count++
			}
		}
		fmt.Printf("count %v of nmber is %v\n", a[i], count)
		if count > (len(a))/2 {
			return a[i], count
		}
	}
	return -1, -1

}

//
//func majoptimize(a []int) (int, int) {
//	count := 0
//	element := math.MinInt
//
//	for i := 0; i < len(a)-1; i++ {
//		if element == a[i] {
//			count++
//		} else {
//			count--
//		}
//		if count == 0 {
//			count = 1
//			element = a[i]
//		}
//	}
//	if count > ((len(a) - 1) / 2) {
//		return element, count
//	}
//	return -1, -1
//}
//
//func main() {
//	// fmt.Println(majoritybrute([]int{1, 4, 4, 4, 4, 4, 4, 4}))
//	fmt.Println(majoptimize([]int{1, 4, 4, 5, 5, 5, 5, 5}))
//}

type Main struct{}

func (m *Main) findCandidate(arr []int, n int) int {
	index := 0
	count := 1
	for i := 1; i < n; i++ {
		if arr[index] == arr[i] {
			count++
		} else {
			count--
		}
		if count == 0 {
			index = i
			count = 1
		}
	}
	fmt.Printf("count %v of nmber is %v\n", arr[index], count)
	return arr[index]
}

func (m *Main) isMajority(arr []int, n int, majorityElement int) bool {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == majorityElement {
			count++
		}
	}
	if count > n/2 {
		return true
	} else {
		return false
	}
}

func (m *Main) printMajority(arr []int, n int) {
	majorityElement := m.findCandidate(arr, n)
	if m.isMajority(arr, n, majorityElement) {
		fmt.Println(majorityElement)
	} else {
		fmt.Println("No Majority Element Found")
	}
}

func main() {
	majorityElement := &Main{}
	arr := []int{9, 4, 3, 9, 9, 9, 9, 9, 8}
	n := len(arr)
	majorityElement.printMajority(arr, n)
	//fmt.Print(majoritybrute(arr))
}
