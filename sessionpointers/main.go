package main

import (
	"errors"
	"fmt"
)

var (
	NOtSupportted  = errors.New("")
	NOtSupportted1 = errors.New("")
	NOtSupportted2 = errors.New("")
	NOtSupportted3 = errors.New("")
)
var NoNameError = errors.New("student doesnt contain name")

type Livingbeing interface {
	Walk()
}

type Person interface {
	Livingbeing
	GetFullName() string
}

type Person1 interface {
	Person
	GetId() int
}

type Student1 struct {
	s           Student
	CollegeName string
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

func (s Student) Walk() string {
	return "hello"
}

func (s Employee) GetFullName() string {
	return s.firstName + s.lastName
}

func main() {
	s := Student{
		id:       10,
		lastName: "def",
		class:    "10",
	}

	//_, str := modifyFirstNameObject(s)
	if name, err := getName(s); err != nil {
		fmt.Printf("error is %v\n", err)
		fmt.Printf("student has no name")
	} else {
		fmt.Printf("stduent name is %v", name)
	}

	//name, err := getName(s)
	////fmt.Printf("name is %v\n", name)
	//if err != nil {
	//	fmt.Printf("error is %v\n", err)
	//	fmt.Printf("student has no name")
	//} else {
	//	fmt.Printf("stduent name is %v", name)
	//}

	//if errors.Is(err, NoNameError) {
	//	fmt.Printf("there is no name")
	//} else {
	//	fmt.Printf("no error")
	//}
}

//
//func readFullName(p Person) {
//	fmt.Printf("person full name is:%v\n", p.GetFullName())
//}
//
//func modifyFirstName(s *Student) *Student {
//	fmt.Printf("address of stduent %p\n", s)
//	s.firstName = "modify"
//	return s
//}
//
//func modifyFirstNameObject(s Student) (Student, string) {
//	fmt.Printf("address of stduent1 %p\n", &s)
//	s.firstName = "modify"
//	return s, "hello"
//}

func getName(s Student) (string, error) {
	lastName := "test"
	defer func(lastname string) {
		fmt.Printf("name of the student %v\n", lastname)
	}(lastName)

	//extra code

	defer fmt.Printf("executing defer1\n")
	if s.firstName == "abc" {
		return s.firstName, nil
	} else {
		lastName = "defer function"
		defer func(lastname string) {
			fmt.Printf("name of the student %v\n", lastname)
		}(lastName)
		return "", NoNameError
	}
}
