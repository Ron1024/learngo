package main

import "fmt"

func main() {
	a := 10
	b := 20
	a, b = change(a, b)

	fmt.Println(a)
	fmt.Println(b)
}

func change(a int, b int) (int, int) {
	return b, a
}
