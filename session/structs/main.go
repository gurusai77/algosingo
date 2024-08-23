package main

import "fmt"

type Person interface {
	GetFullName() string
	GetID() int
}

type Student struct {
	Id        int
	FirstName string
	LastName  string
	Class     string
}

type EnggStudent struct {
	Student Student
	Branch  string
}

type SchoolStudent struct {
	Student
	SchoolName string
}

func (s Student) GetID() int {
	return s.Id
}

type Employee struct {
	id        int
	firstName string
	lastName  string
}

func (e Employee) GetID() int {
	return e.id
}

type Worker struct {
	id        int
	firstName string
	lastName  string
}

func (w Worker) GetID() int {
	return w.id
}

func (w Worker) GetFullName() string {
	return w.firstName + w.lastName
}

func (s Student) GetFullName() string {
	return s.FirstName + s.LastName
}

func (e Employee) GetFullName() string {
	return e.firstName + e.lastName
}

func NewStudent(fname, lname, class string, id int) Student {
	return Student{
		Id:        id,
		FirstName: fname,
		LastName:  lname,
		Class:     class,
	}
}

func main() {
	s := Student{
		Id:        10,
		FirstName: "abc",
		LastName:  "def",
		Class:     "10",
	}
	//s1 := NewStudent("first", "last", "seven", 10)
	//fmt.Printf("student name is %v \n", s1.GetFullName())
	//fmt.Printf("student name is %v \n", s.GetFullName())
	//e := Employee{
	//	id:        10,
	//	firstName: "emplo",
	//	lastName:  "yee",
	//}
	//readFullName(s)
	//readFullName(e)
	//w := Worker{
	//	id:        10,
	//	firstName: "worker",
	//	lastName:  "-worker",
	//}
	//readFullName(w)

	es := EnggStudent{
		Student: Student{
			Id:        10,
			FirstName: "abc",
			LastName:  "def",
			Class:     "10",
		},
		Branch: "IT",
	}
	ss := SchoolStudent{
		Student:    s,
		SchoolName: "IT",
	}
	fmt.Printf("id of student is %v \n", es.Student.Id)
	fmt.Printf("student obj is %v \n", es.Branch)

	fmt.Printf("id of school student is %v \n", ss.Student.Id)
}

func readFullName(p Person) {
	fmt.Printf("person full name is:%v\n", p.GetFullName())
	switch p.(type) {
	case Student:
		fmt.Printf("this is student \n")
	case Employee:
		fmt.Printf("this is employee \n")
	case Worker:
		fmt.Printf("this is worker \n")
	default:
		fmt.Printf("this is default \n")
	}
}
