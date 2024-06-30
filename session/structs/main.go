package main

import "fmt"

type Person interface {
	GetFullName() string
	// GetID() int
}

type Student struct {
	id        int
	firstName string
	lastName  string
	class     string
}

type Employee struct {
	id        int
	firstName string
	lastName  string
}

func (s Student) GetFullName() string {
	return s.firstName + s.lastName
}

func (s Employee) GetFullName() string {
	return s.firstName + s.lastName
}

func (s Student) GetName() string {
	return s.firstName
}

func (s Student) GetId() int {
	return s.id
}

func (s Student) GetClass() string {
	return s.class
}

func main() {
	s := Student{
		id:        10,
		firstName: "abc",
		lastName:  "def",
		class:     "10",
	}
	e := Employee{
		id:        10,
		firstName: "emplo",
		lastName:  "yee",
	}
	readFullName(s)
	readFullName(e)
}

func readFullName(p Person) {
	fmt.Printf("person full name is:%v\n", p.GetFullName())
	switch p.(type) {
	case Student:
		fmt.Printf("this is student")
	case Employee:
		fmt.Printf("this is employee")
	default:
		fmt.Printf("this is default")
	}
}
