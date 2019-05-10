package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//Student 学生结构体
type Student struct {
	Name string
	ID   string
	Age  int
}

//Book 书籍结构体
type Book struct {
	Name   string
	Author string
}

//StudentArray 定义一个学生的切片别名
type StudentArray []Student

//Len 计算长度
func (p StudentArray) Len() int {
	return len(p)
}

//Less 比较大小
func (p StudentArray) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

//Swap 交换位置
func (p StudentArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var stus StudentArray
	//随机生成10个学生
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			ID:   fmt.Sprintf("510%d", rand.Int()),
			Age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	//打印学生的信息
	for _, v := range stus {
		fmt.Println(v)
	}

	fmt.Println()

	fmt.Println()
	sort.Sort(stus)
	for _, v := range stus {
		fmt.Println(v)
	}
}
