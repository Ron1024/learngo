package main

import "fmt"

func main() {
	a := 100
	b := 200
	swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
