package main

import "fmt"

type Test interface {
	Print()
	Sleep()
}

type People struct {
	name string
	age  int
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

func (stu *Student) Sleep() {
	fmt.Println(stu.name + " is sleeping")
}

func (p *People) Print() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

func (p *People) Sleep() {
	fmt.Println("sleeping ", p.name)
}

func main() {
	var t Test

	stu := Student{
		name:  "Parker",
		age:   10,
		score: 20,
	}

	pp := People{
		name: "lili",
		age:  30,
	}

	t = &stu

	t.Print()

	t.Sleep()

	t = &pp
	t.Print()
	t.Sleep()

}
