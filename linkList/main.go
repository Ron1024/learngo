package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func main() {
	var head Student
	head.Name = "Parker"
	head.Age = 28
	head.Score = 99.9

	var stu1 Student
	stu1.Name = "Nina"
	stu1.Age = 30
	stu1.Score = 89.9

	head.next = &stu1

	var stu2 Student
	stu2.Name = "lili"
	stu2.Age = 30
	stu2.Score = 89.9

	head.next = &stu1

	stu1.next = &stu2

	var p *Student = &head

	trans(p)
}

func trans(p *Student) {

	for p != nil {

		fmt.Println(*p)

		p = (*p).next

	}

}
