package main

import "fmt"

func main() {
	num := calc(50)
	fmt.Println(num)
}

func calc(n int64) int64 {
	if n <= 0 {
		return 1
	}

	return n * calc(n-1)
}
