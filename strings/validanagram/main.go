package main

import "fmt"

//Time complexity: O(N)
//
//The time complexity of the above code is O(n), where n is the length of the longer string between s and t. This is because the code iterates over each character in both strings only once, using two separate loops.
//The first loop counts the occurrences of characters in string s, and the second loop subtracts the occurrences of characters in string t. Finally, there is a third loop that checks if any count in the alphabet array is nonzero, indicating a mismatch in character counts between the two strings.
//
//Space complexity: O(1) or Constant
//
//The space complexity of the code is O(1) or constant because uses an integer array, alphabet, of size 26 to store the counts of characters in the English alphabet.
//Regardless of the length of the input strings, the alphabet array remains a fixed size because it represents a fixed set of characters (lowercase English letters).
//Therefore, the space required by the array is constant and does not depend on the input size.

func main() {
	s := "anagram"
	t := "naagarm"
	fmt.Println(s)
	fmt.Println(t)
	var alpha = make([]int, 26)
	for i := 0; i < len(s); i++ {
		alpha[s[i]-'a']++
	}
	fmt.Println(alpha)
	for i := 0; i < len(s); i++ {
		alpha[t[i]-'a']--
	}
	fmt.Println(alpha)
	for i := 0; i < len(alpha); i++ {
		if alpha[i] != 0 {
			fmt.Println("string is not a valid anagram")
		}
	}
	fmt.Println("string is a valid anagram")
}
