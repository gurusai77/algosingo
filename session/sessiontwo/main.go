package main

import (
	"fmt"
)

var i = 10

const J = 20

//func init() {
//	fmt.Printf("this is init function")
//	fmt.Printf("value of k :%v \n", k)
//}
//
//func init() {
//	fmt.Printf("this is init function2")
//}

var k = 20

func main() {
	var m = make(map[string]bool)
	//var m = make(map[string]string)
	//var m = make(map[int]int)
	//var m = make(map[string]float64)

	m["key"] = true
	m["key1"] = true
	m["key2"] = false
	//for k, v := range m {
	//	fmt.Printf("key is %v and values %v \n", k, v)
	//}
	//_, _, k := returntwovalues()
	//
	//fmt.Printf("%v", k)
	// receiveMaps(m)
	//loopMap(m)
	//
	//	var a = [0]int{}
	//	for index, value := range a {
	//		fmt.Printf("index is %v, value is %v \n", index, value)
	//	}
	//	loopString(10, "s", "s2", "s3s")
	//}
	//
	//func receiveMaps(m map[string]bool) {
	//	fmt.Printf("print key value: %v\n", m["key2"])
	//	delete(m, "key1")
	//	fmt.Print(m)

	//loopString(1)
	//loopString(1, "s2", "s3", "s4", "s5")

	switchexample("def")
}

func returntwovalues() (int, int, int) {
	return 1, 2, 4
}

func loopMap(m map[string]bool) {
	delete(m, "key")
	delete(m, "key1")
	for k, v := range m {
		fmt.Printf("key is %v, value is %v \n", k, v)
	}
}

func loopString(i int, s ...string) (int, string) {
	for k, v := range s {
		fmt.Printf("index is %v value is %v \n", k, v)
	}
	return 10, "string_value"
}

func switchexample(i string) {
	switch i {
	case "december":
		fmt.Printf("type of i is %T \n", i)
		fmt.Printf("value of i is %v \n ", i)
	case "november":
		fmt.Printf("type of i is %T \n", i)
		fmt.Printf("value of i is %v \n", i)
	default:
		fmt.Printf("def type of i is %T \n", i)
		fmt.Printf("def value of i is %v \n", i)
	}
}
