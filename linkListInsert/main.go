package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student

	//p	func toString() strin{
	//	return "{Name = "+Name+",Age="+Age+",Score="+Score+"}"
	//	}
}

func trans(p *Student) {

	for p != nil {
		fmt.Println(*p)
		p = (*p).next
	}

}
func main() {
	var stu *Student = new(Student)
	stu.Name = "Parker"
	stu.Age = 29
	stu.Score = 99.9
	insertHead(&stu)
	trans(stu)
}

func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 1000; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = &stu
		tail = &stu
	}

}

//链表前方插入
func insertHead(p **Student) {

	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = &stu

	}
}
