package main

import (
	"fmt"
	"learngo/channle/add"
)

func main() {

	var pipe chan int
	pipe = make(chan int, 2)
	pipe <- 1
	pipe <- 5

	fmt.Println(add.Add(pipe))

	fmt.Println(add.Clac(4, 2))
}
