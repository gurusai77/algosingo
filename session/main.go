package main

import "fmt"

//var s = 10

//func init() {
//	fmt.Println("init")
//	fmt.Println(s)
//	fmt.Println(s1)
//}

var s1 = 10

//func main() {
//	var arr = [4]int{1, 2, 3, 4}
//	fmt.Printf("cap :%v \n", cap(arr))
//	fmt.Printf("length :%v \n", len(arr))
//	fmt.Printf("size of arr :%v \n", unsafe.Sizeof(arr))
//	var s = []int{1, 2, 3, 4}
//	fmt.Printf("cap :%v \n", cap(s))
//	fmt.Printf("length :%v \n", len(s))
//	fmt.Printf("type :%T \n", s)
//	fmt.Printf("size of slice :%v \n", unsafe.Sizeof(s))
//
//	var k = make([]int, 5)
//	j := slicePointer(k)
//	fmt.Println(j)
//	fmt.Println(&k[0])
//	if j == &k[0] {
//		fmt.Println("true")
//	}
//}

//	func slicePointer(s []int) *int {
//		return &s[0]
//	}
const SESSIONVALUE = 1

//
//var (
//	Cookie      = ""
//	CookieValue = "username"
//	SessionId   = ""
//)

func main() {
	var i int = 5
	//s := "hello"
	fmt.Printf("variable value of i :%v \n", i)
	fmt.Printf("variable value of i :%T \n", i)
	for i <= 10 {
		fmt.Printf("variable in loop value of i :%v \n", i)
		fmt.Printf("i value is: %v\n", i)
		i++
	}
	for i := 0; i <= 10; i++ {
		fmt.Printf("variable in loop value of i :%v \n", i)
		fmt.Printf("i value is: %v\n", i)
	}

	// arrays and slices
	a := [5]string{"1", "2", "2", "4", "5"}
	fmt.Printf("array :%v \n", a)

	b := []string{"1", "2", "2", "4", "5"}
	c := make([]float64, 10)
	d := &b
	fmt.Printf("slice :%v", b)
	fmt.Printf("slice :%v", c)
	fmt.Printf("slice :%p", b)
	fmt.Printf("slice address:%p", &d)

	j := 10
	fmt.Printf("variable value of j :%v \n", j)
	fmt.Printf("variable value of j :%T \n", j)

	k, l := 10, true
	fmt.Printf("variable value of k :%v \n", k)
	fmt.Printf("variable value of k :%d \n", k)
	fmt.Printf("variable value of l :%T \n", l)

	fmt.Printf("constant value of session :%v \n", SESSIONVALUE)
	var m rune = 'a'
	fmt.Printf("value of a :%v \n", m)
	fmt.Printf("type of a :%T \n", m)

	var n string = "sample"
	fmt.Printf("value  of b :%v \n", n)
	fmt.Printf("value  of b :%T \n", n)

	if i == j {
		fmt.Print("tue")
	}

}
