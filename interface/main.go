package main

import "fmt"

type Test interface {
	Print()
}

type Student struct {
	name  string
	age   int
	score int
}

func (stu *Student) Print() {
	fmt.Println(stu.name)
	fmt.Println(stu.age)
	fmt.Println(stu.score)
}

func main() {
	var t Test

	var stu Student = Student{
		name:  "Parker",
		age:   10,
		score: 20,
	}

	t = &stu

	t.Print()

}
