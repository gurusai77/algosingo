package main

import (
	"fmt"
)

var i = 10

const J = 20

func init() {
	fmt.Printf("this is init function")
	fmt.Printf("value of k :%v \n", k)
}

func init() {
	fmt.Printf("this is init function2")
}

var k = 20

func main() {
	var m = make(map[string]bool)
	m["key"] = true
	m["key1"] = true
	// receiveMaps(m)
	loopMap(m)

	var a = [0]int{}
	for index, value := range a {
		fmt.Printf("index is %v, value is %v \n", index, value)
	}
	loopString(10, "s", "s2", "s3s")
}

func receiveMaps(m map[string]bool) {
	fmt.Printf("print key value: %v\n", m["key2"])
	delete(m, "key1")
	fmt.Print(m)
}

func loopMap(m map[string]bool) {
	delete(m, "key")
	delete(m, "key1")
	for k, v := range m {
		fmt.Printf("key is %v, value is %v \n", k, v)
	}
}

func loopString(i int, s ...string) (int, string) {
	for _, v := range s {
		fmt.Printf("value is %v \n", v)
	}
	return 10, "string_value"
}
