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
	// linkListDelete(stu)
	var newNode *Student = new(Student)
	newNode.Name = "nima"
	newNode.Age = 20
	newNode.Score = 0.1
	addNode(stu, newNode)
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

func linkListDelete(p *Student) {
	var prew *Student = p
	for p != nil {
		if p.Name == "stu6" {
			prew.next = p.next
			break
		}
		prew = p
		p = p.next
	}
}

func addNode(p, newNode *Student) {
	for p != nil {
		if p.Name == "stu5" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}
