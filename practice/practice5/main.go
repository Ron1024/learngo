package main

import "fmt"

func main() {
	a := 10
	var b *int
	b = &a
	fmt.Println(a)
	fmt.Println(*b)
}
