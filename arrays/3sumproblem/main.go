package main

import "fmt"

// brute force, o(n3) and space complexity as we are using extra array
func tripletsSum(a []int, sum int) (l [][]int) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n-1; k++ {
				if a[i]+a[j]+a[k] == sum {
					l = append(l, []int{a[i], a[j], a[k]})
				}
			}
		}
	}
	return l
}

// better approach using pointers 0(n2)
func tripSum(a []int, reqSum int) (l [][]int) {
	n := len(a)
	for i := 0; i < n; i++ {
		// handle duplicates
		if i != 0 && a[i] == a[i-1] {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := a[i] + a[j] + a[k]
			if sum < reqSum {
				j++
			} else if sum > reqSum {
				k--
			} else {
				l = append(l, []int{a[i], a[j], a[k]})
				j++
				k--
				// handle duplicates
				for j < k && a[j] == a[j-1] {
					j++
				}
				for j < k && a[k] == a[k+1] {
					k--
				}
			}
		}
	}
	return l
}

func main() {
	// pass a sorted array
	a := tripletsSum([]int{1, 2, 3, 4, 5, 6}, 6)
	fmt.Print(a)
	b := tripSum([]int{-4, -1, -1, 0, 1, 2}, 0)
	fmt.Print(b)
	b = tripSum([]int{1, 2, 2, 3, 4, 5, 6}, 6)
	fmt.Print(b)
}
