package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func myAtoi(s string) int {
	r, _ := regexp.Compile(`[A-Za-z ]+`)
	s = r.ReplaceAllString(s, "")
	i, _ := strconv.ParseInt(s, 10, 0)
	return int(i)
}

func main() {
	fmt.Println(myAtoi("words and987"))
}
