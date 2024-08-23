package main

import "fmt"

// better approach using pointers o(n3)
func fourSumProblem(a []int, reqSum int) (store [][]int) {
	n := len(a)
	for i := 0; i < n; i++ {
		if i != 0 && a[i] == a[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j != 0 && a[j] == a[j-1] {
				continue
			}
			k, l := j+1, n-1
			for k < l {
				sum := a[i] + a[j] + a[k] + a[l]
				if sum < reqSum {
					k++
				} else if sum > reqSum {
					l--
				} else {
					store = append(store, []int{a[i], a[j], a[k], a[l]})
					k++
					l--
					for k < l && a[k] == a[k-1] {
						k++
					}
					for k < l && a[l] == a[l+1] {
						l--
					}
				}
			}
		}
	}
	return store
}

func main() {
	// pass a sorted array
	a := fourSumProblem([]int{1, 2, 3, 4, 5, 6, 7, 8}, 22)
	fmt.Print(a)
}
