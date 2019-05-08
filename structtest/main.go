package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var stu Student

	stu.Age = 10
	stu.Name = "Parker"

	fmt.Println(stu)

	var stu1 *Student = &Student{
		Name: "Lili",
		Age:  30,
	}

	fmt.Println(stu1.Name)

	var stu2 = Student{
		Name: "Nina",
		Age:  50,
	}

	fmt.Println(stu2.Name)
}
