package main

import "fmt"

func substring(s string) int {
	left, right := 0, 0
	maxLen, startIndex, lastIndex := 0, 0, 0
	visited := [256]bool{}
	for right < len(s) {
		if visited[s[right]] {
			for visited[s[right]] {
				visited[s[left]] = false
				left++
			}
		}
		visited[s[right]] = true
		if right-left+1 > maxLen {
			startIndex = left
			lastIndex = right
		}
		maxLen = max(maxLen, (right-left)+1)
		right++
	}
	fmt.Printf("maxLen is %v \n", maxLen)
	fmt.Printf("max string indexes %v, %v \n", startIndex, lastIndex)
	fmt.Printf("max string is %v \n", s[startIndex:lastIndex+1])
	return maxLen
}

func longestUniqueSubsttr(str string) int {
	if len(str) == 0 {
		return 0
	}

	if len(str) == 1 {
		return 1
	}

	maxLength := 0
	visited := make([]bool, 256)

	left := 0
	right := 0
	for right < len(str) {
		if visited[str[right]] {
			for visited[str[right]] {
				visited[str[left]] = false
				left++
			}
		}
		visited[str[right]] = true

		maxLength = max(maxLength, (right - left + 1))
		right++
	}
	fmt.Printf("maxlen: %v \n", maxLength)
	return maxLength
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func lengthOfLongestSubstring(s string) int {
	var occurrencesMap = map[rune]int{}
	maxCount := 0
	// pointer to where the sequence starts
	current := 0
	for index, runeValue := range s {
		// 0 is the default value
		// If a char has been seen before and the index of the duplicate is after in the word than our last seen duplicate, update the index of the last seen duplicate and start over
		if occurrencesMap[runeValue] != 0 && occurrencesMap[runeValue] >= current {
			current = occurrencesMap[runeValue]
		}
		// If the current index minus the pointer to where the last duplicate was seen is bigger then the last amx, then we have a longer sequence
		currentLength := index + 1 - current
		if currentLength >= maxCount {
			maxCount = currentLength
		}
		// we store the index as i + 1 so in the map if it is 0 then it is null
		occurrencesMap[runeValue] = index + 1
	}
	fmt.Printf("string is: %v \n", maxCount)
	return maxCount
}

func main() {
	lengthOfLongestSubstring("geeksforgeeks")
	longestUniqueSubsttr("geeksforgeeks")
	substring("geeksforgeeks")
}
