package main

import (
	"errors"
	"log"
	"os"
)

func reverseString(s string) string {
	i := 0
	j := len(s) - 1
	rns := []rune(s)
	for i <= len(rns)/2 && i < j {
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
	//fmt.Println(reverseString("stephen"))
	//fmt.Println(reverse("stephen"))
	output := "./output"
	path := output
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
