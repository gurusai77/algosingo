package main

import (
	"fmt"
	"sync"
)

func f1(wg *sync.WaitGroup, c chan int, n *int, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	*n++
	c <- *n
	wg.Done()
}

func f2(wg *sync.WaitGroup, c chan int, n *int, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	*n--
	c <- *n
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	c := make(chan int)

	var m sync.Mutex
	n := 0

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go f1(&wg, c, &n, &m)
		n = <-c
		fmt.Println(n)
		go f2(&wg, c, &n, &m)
		n = <-c
		fmt.Println(n)
	}
	fmt.Println(n)
	wg.Wait()
	close(c)
	c <- 2
}
