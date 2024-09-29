package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "IceCreAm"

	runeSlice := []rune(s)
	fmt.Println(runeSlice)
	low, high := 0, len(runeSlice)-1
	vowels := "aeiouAEIOU"
	for low < high {

		if strings.Contains(vowels, string(runeSlice[low])) &&
			strings.Contains(vowels, string(runeSlice[high])) {
			runeSlice[low], runeSlice[high] = runeSlice[high], runeSlice[low]
			low++
			high--
		} else if !strings.Contains(vowels, string(runeSlice[low])) {
			low++

		} else if !strings.Contains(vowels, string(runeSlice[high])) {
			high--
		}

	}
	fmt.Println()
	fmt.Println(string(runeSlice))
}
