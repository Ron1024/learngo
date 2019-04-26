package main

import (
	"fmt"
	add "learngo/practice/practice2/calc"
)

const (
	a = iota
	b
	c
	d
	e
)

func main() {
	fmt.Println(add.Name)
	fmt.Println(add.Age)

	fmt.Println("b = ", b)

}
