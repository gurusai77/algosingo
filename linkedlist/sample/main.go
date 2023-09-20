package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func (l *List) Remove(d interface{}) {
	p := l.head
	for p.next != nil {
		if p.info == d {
			p.info = p.next.info
			p.next = p.next.next
			break
		} else {
			p = p.next
		}

	}
}

func (l *List) ConvertToCircular() {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = l.head
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
}

func (l *List) ShowCircular() {
	p := l.head
	for {
		if p.next == l.head {
			fmt.Printf("-> %v ", p.info)
			break
		}
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
}

func main() {
	sl := List{}
	for i := 0; i < 5; i++ {
		sl.Insert(rand.Intn(100))
	}
	Show(&sl)
	sl.Insert(23)
	sl.Insert(24)
	sl.Insert(56)
	sl.Insert(57)
	fmt.Println()
	Show(&sl)
	sl.Remove(23)
	fmt.Println()
	Show(&sl)
	sl.ConvertToCircular()
	fmt.Println()
	sl.ShowCircular()
}
