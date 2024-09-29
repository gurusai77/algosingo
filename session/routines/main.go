package main

import (
	"fmt"
	"sync"
	"time"
)

func sayhello(c chan string) {
	fmt.Printf("hello \n")
	c <- "hello"
}

func sayhello1() {
	fmt.Printf("hello1 \n")
}

func sayhello2(wg *sync.WaitGroup) {
	fmt.Printf("hello2 \n")
	fmt.Printf("address of wg in hello2 %p \n", &wg)
	wg.Done()
}

func main() {
	//c := make(chan string)
	fmt.Printf("time is %v\n", time.Now())
	fmt.Printf("main function \n")
	//go sayhello(c)
	//str := <-c
	var wg sync.WaitGroup
	fmt.Printf("address of wg in main %p \n", &wg)
	//go sayhello1()
	wg.Add(1)
	go sayhello2(&wg)
	wg.Wait()

	//fmt.Printf("received from channel:%v \n", str)
	//go func() {
	//	fmt.Printf("anonymous func \n")
	//}()
	// time.Sleep(time.Second)
	fmt.Printf("completed main function \n")
	fmt.Printf("time is %v\n", time.Now())
}

//func main() {
//	var wg sync.WaitGroup
//
//	for i := 0; i <= 2; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			sayhello1()
//		}()
//	}
//	wg.Wait()
//}

//func main() {
//	messages := make(chan int)
//	var wg sync.WaitGroup
//	var result []int
//
//	// you can also add these one at
//	// a time if you need to
//
//	//wg.Add(1)
//	//go func() {
//	//	defer wg.Done()
//	//	time.Sleep(time.Second * 1)
//	//	messages <- 1
//	//}()
//	//wg.Add(1)
//	//go func() {
//	//	defer wg.Done()
//	//	time.Sleep(time.Second * 1)
//	//	messages <- 2
//	//}()
//	//wg.Add(1)
//	//go func() {
//	//	defer wg.Done()
//	//	time.Sleep(time.Second * 1)
//	//	messages <- 3
//	//}()
//
//	for i := 0; i <= 2; i++ {
//		wg.Add(1)
//		//j := i
//		go func(i int) {
//			defer wg.Done()
//			messages <- i
//		}(i)
//	}
//
//	// this goroutine added to signal end of data stream
//	// by closing messages channel
//	go func() {
//		wg.Wait()
//		close(messages)
//	}()
//
//	// if you need this to happen inside a go routine,
//	// this channel is used for signalling end of the work,
//	// also another sync.WaitGroup could be used, but for just one
//	// goroutine, a single channel as a signal makes sense (there is no
//	// groups)
//
//	for i := range messages {
//		fmt.Println(i)
//		result = append(result, i)
//	}
//
//	fmt.Println(result)
//}
