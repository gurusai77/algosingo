package main

import "fmt"

func besttime(a []int) int {
	minPrice := a[0]
	maxProfit := 0
	for i := 1; i < len(a); i++ {
		if minPrice > a[i] {
			minPrice = a[i]
		}
		if maxProfit < a[i]-minPrice {
			maxProfit = a[i] - minPrice
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(besttime([]int{7, 1, 7, 3, 8, 5, 9}))
}
