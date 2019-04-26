package main

import "fmt"

func main() {
	group(10)
}

func group(a int) {

	var b int
	var c int
	for i := 0; i < a; i++ {
		b = i
		c = a - b

		fmt.Printf("%d+%d = %d\n", b, c, a)
	}
}
