package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (stu *Student) init(name string, age int) {
	stu.name = name
	stu.age = age
	fmt.Println(stu)
}

func main() {
	var stu Student
	stu.init("Parker", 20)

	stu1 := stu
	fmt.Println(stu1)
}
