package main

import "fmt"

func movezeros(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr
}

func main() {
	arr := []int{1, 2, 0, 0, 2, 1, 0, 1, 1, -1, -10}
	fmt.Println(movezeros(arr))
}
