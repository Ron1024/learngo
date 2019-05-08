package main

import (
	"fmt"
	"strings"
)

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))

	mak := makeSuffix("http")
	str1 := mak("p")
	fmt.Println(str1)
}

func Adder() func(int) int {
	var x int
	return func(a int) int {
		x += a
		return x
	}
}

func makeSuffix(name string) func(str string) string {

	return func(str string) string {
		if strings.HasSuffix(name, str) == false {
			return name + str
		}
		return name
	}
}
