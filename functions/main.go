package main

import "fmt"

func main() {
	c := add
	sum := c(200, 400)
	fmt.Println(sum)

	fmt.Println(operator(add, 100, 200))
}

func add(a, b int) int {
	return a + b
}

type addFunc func(int, int) int

func operator(op addFunc, a, b int) int {
	return op(a, b)
}
