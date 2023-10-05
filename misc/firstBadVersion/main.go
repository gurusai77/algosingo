package main

import "fmt"

func isBadVersion(version int, firstBad int) bool {
	if version >= firstBad {
		return true
	}
	return false
}

func main() {
	// badVersion := 1
	size := 5
	bad := predictbadversion(size, 3)
	if bad == size {
		fmt.Println("all are good versions")
	} else {
		fmt.Println(bad)
	}
}

func predictbadversion(size int, badVersion int) int {
	start := 1
	end := size

	for start < end {
		mid := (start + end) / 2
		if isBadVersion(mid, badVersion) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
