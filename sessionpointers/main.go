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
	CannotTalk     = errors.New("animals cannot talk")
)
var NoNameError = errors.New("student doesnt contain name")

type LivingBeing interface {
	Walk() bool
	Talk() (bool, error)
}

type HumanBeing interface {
	LivingBeing
	GetFullName() (*string, error)
}

type Animal struct {
	Name string
}

func (a Animal) Walk() bool {
	return true
}

func (a Animal) Talk() (bool, error) {
	return false, CannotTalk
}

type Student struct {
	Name      string
	firstName string
	lastName  string
}

func (s Student) GetFullName() (*string, error) {
	defer func() {
	}()
	if s.Name == "" {
		return nil, NoNameError
	}
	defer fmt.Printf("student has name defined in struct \n")
	name := s.Name
	return &name, nil
}

func (s Student) Walk() bool {
	return true
}

func (s Student) Talk() (bool, error) {
	return true, nil
}

func main() {
	s := Student{
		Name:      "",
		firstName: "",
		lastName:  "def",
	}

	a := Animal{
		Name: "Lion",
	}

	fmt.Printf("can walk: %v \n", a.Walk())

	if talk, err := a.Talk(); err != nil {
		fmt.Printf("cannot talk: %v \n", talk)
		fmt.Printf("animals cannot talk\n")
	} else {
		fmt.Printf("animals can talk")
	}

	name, err := s.GetFullName()
	if err == nil {
		fmt.Printf("name of student is: %v \n", name)
	} else {
		fmt.Printf("student doesnt contain name \n")
	}
	s.Name = "hello"
	name, err = s.GetFullName()
	if err == nil {
		fmt.Printf("name of student is: %v \n", *name)
	} else {
		fmt.Printf("student doesnt contain name \n")
	}

	_, deferErr := getName(s)

	if deferErr != nil {
		fmt.Print(s.lastName)
	} else {
		fmt.Print(s.lastName)
	}

	//modifyFirstNameObject(s)
	//if name, err := getName(s); err != nil {
	//	fmt.Printf("error is %v\n", err)
	//	fmt.Printf("student has no name")
	//} else {
	//	fmt.Printf("stduent name is %v", name)
	//}
	//
	////name, err := getName(s)
	//////fmt.Printf("name is %v\n", name)
	////if err != nil {
	////	fmt.Printf("error is %v\n", err)
	////	fmt.Printf("student has no name")
	////} else {
	////	fmt.Printf("stduent name is %v", name)
	////}
	//
	////if errors.Is(err, NoNameError) {
	////	fmt.Printf("there is no name")
	////} else {
	////	fmt.Printf("no error")
	////}
}

//	func readFullName(p Person) {
//		fmt.Printf("person full name is:%v\n", p.GetFullName())
//	}
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
	s.lastName = "test"
	defer func(s *Student) {
		fmt.Printf("name of the student %v\n", s.lastName)
	}(&s)

	//extra code

	defer fmt.Printf("executing defer1\n")
	if s.firstName == "abc" {
		return s.firstName, nil
	} else {
		//s.lastName = "defer function"
		defer func(s *Student) {
			fmt.Printf("name of the student %v\n", s.lastName)
		}(&s)
		return "", NoNameError
	}
}
