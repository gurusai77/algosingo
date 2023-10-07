package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	for i := low; i < high; i++ {
		if arr[i] < pivot {
			arr[i], arr[low] = arr[low], arr[i]
			i++
		}
	}
	arr[low], arr[high] = pivot, arr[low]
	return arr, low
}

func quickS(a []int, low int, high int) []int {
	if low < high {
		a, p := partition(a, low, high)
		fmt.Printf("new low high and p: %v, %v, %v \n", low, p, high)
		quickS(a, low, p-1)
		quickS(a, p+1, high)
	}
	return a
}

func partitionTwo(a []int, low int, high int) ([]int, int) {

	fmt.Printf("new low and high: %v, %v \n", low, high)
	pivot := a[low]
	fmt.Printf("new pivot: %v \n", pivot)
	i := low
	j := high
	for i < j {
		for a[i] <= pivot && i <= high-1 {
			i++
		}
		for a[j] > pivot && j >= low+1 {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[j], a[low] = a[low], a[j]
	fmt.Printf("end of partition: %v \n", a)
	return a, j
}

func quickSort(a []int, low int, high int) []int {
	if low < high {
		a, p := partitionTwo(a, low, high)
		fmt.Printf("new low high and p: %v, %v, %v \n", low, p, high)
		quickSort(a, low, p-1)
		quickSort(a, p+1, high)
	}
	return a
}

func partitionD(a []int, low, high int) ([]int, int) {
	fmt.Printf("new low and high: %v, %v \n", low, high)
	pivot := a[low]
	fmt.Printf("new pivot: %v \n", pivot)
	i := low
	j := high
	for i < j {
		for a[i] >= pivot && i <= high-1 {
			i++
		}
		for a[j] < pivot && j >= low+1 {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[j], a[low] = a[low], a[j]
	fmt.Printf("end of partition: %v \n", a)
	return a, j
}

func quickD(a []int, low int, high int) []int {
	if low < high {
		a, p := partitionD(a, low, high)
		fmt.Printf("new low high and p: %v, %v, %v \n", low, p, high)
		quickD(a, low, p-1)
		quickD(a, p+1, high)
	}
	return a
}

func main() {
	a := []int{24, 9, 19, 27, 18, 10, 11, 28}
	fmt.Println(quickD(a, 0, len(a)-1))
	a = []int{24, 9, 19, 27, 18, 10, 11}
	fmt.Println(quickSort(a, 0, len(a)-1))
}
