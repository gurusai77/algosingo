package main

import "fmt"

func reverseString(s string) string {
	i := 0
	j := len(s) - 1
	rns := []rune(s)
	for i < len(rns)/2 && i < j {
		rns[i], rns[j] = rns[j], rns[i]
		i++
		j--
	}
	return string(rns)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println(reverseString("stephen"))
	fmt.Println(reverse("stephen"))
}
