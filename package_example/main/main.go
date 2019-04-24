package main

import (
	"fmt"
	"learngo/package_example/calc"
)

func main() {
	a := 4

	b := 5
	fmt.Println(calc.Add(a, b))
	fmt.Println(calc.Device(a, b))
}
