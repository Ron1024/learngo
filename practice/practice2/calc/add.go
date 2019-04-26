package add

import (
	"fmt"
	_ "learngo/practice/practice2/test"
)

var (
	Name string
	Age  int
)

//
//func Add() (string, int) {
//	Name = "Parker"
//	Age = 16
//	return Name, Age
//}

func init() {
	Name = "parker"
	Age = 18

	fmt.Println("This is the add package init(), name = ", Name)
	fmt.Println("This is the add package init(), age  = ", Age)
}
