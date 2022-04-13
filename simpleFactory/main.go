package main

import "fmt"

type student struct {
	name string
	age  int
}

func (s *student) gotoSchool() {
	fmt.Println(s.name, " gotoSchool...")
}

func NewStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func main() {
	zhangsan := NewStudent("张三", 3)
	zhangsan.gotoSchool()

	lisi := NewStudent("李四", 4)
	lisi.gotoSchool()
}
