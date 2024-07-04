package main

import (
	"fmt"
	"sync"
)

type single struct {
	name string
}

var lock = &sync.Mutex{}
var s *single
var wg sync.WaitGroup

func main() {

	//wg.Add(1)
	//singletonObj(&wg)
	//wg.Add(1)
	//go singletonObj(&wg)
	//wg.Add(1)
	//go singletonObj(&wg)
	//wg.Wait()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			singletonObj()
		}()
	}
	wg.Wait()
}

func singletonObj() {
	if s == nil {
		lock.Lock()
		defer lock.Unlock()
		if s == nil {
			fmt.Println("creating first instance")
			s = &single{
				name: "test",
			}
			fmt.Println("instance created for first time")
		} else {
			fmt.Println("instance already created")
		}
	} else {
		fmt.Println("instance already created")
	}
}
