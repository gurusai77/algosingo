package main

import (
	"fmt"
	"math"
)

/*

Let's observe the pattern of adding the units digit from 1 to 18 below:

1 -> 1
2 -> 2
3 -> 3
4 -> 4
5 -> 5
6 -> 6
7 -> 7
8 -> 8
9 -> 9
10 -> 1
11 -> 2
12 -> 3
13 -> 4
14 -> 5
15 -> 6
16 -> 7
17 -> 8
18 -> 9

We can observe that the result of adding the units digit appears in a cycle,
which is from 1 to 9 repeating constantly. Therefore, we can simply take the given number modulo 9,
and the result will be the sum of the digits. If the result is 0, then we return 9.

*/

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

func sDigits(num int) int {
	if num < 10 {
		return num
	}

	var result int
	for num > 0 {
		result = result + num%10
		num = num / 10
	}

	return sDigits(result)
}

func reverseDigits(num int) int {
	if num < 10 {
		return num
	}

	var result int
	for num != 0 {
		s := num % 10
		result = result*10 + s
		num = num / 10
	}

	return result
}

func countDigits(a int) float64 {
	return math.Log10(float64(a))
}

func main() {
	fmt.Println(sDigits(345678))
	fmt.Println(addDigits(78))
	fmt.Println(reverseDigits(78))
	fmt.Println(math.Ceil(countDigits(78828391854)))
}
