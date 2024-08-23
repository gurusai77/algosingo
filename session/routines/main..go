package main

import (
	"fmt"
	"time"
)

func sayhello(c chan string) {
	fmt.Printf("hello \n")
	c <- "hello"
}

func sayhello1() {
	fmt.Printf("hello1 \n")
}

func sayhello2() {
	fmt.Printf("hello2 \n")
}

func main() {
	c := make(chan string)
	fmt.Printf("time is %v\n", time.Now())
	fmt.Printf("main function \n")
	go sayhello(c)
	str := <-c
	go sayhello1()
	go sayhello2()

	fmt.Printf("received from channel:%v \n", str)
	//go func() {
	//	fmt.Printf("anonymous func \n")
	//}()
	time.Sleep(time.Second)
	fmt.Printf("completed main function \n")
	fmt.Printf("time is %v\n", time.Now())
}
