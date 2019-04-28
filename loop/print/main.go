package main

import "fmt"

func main() {
	printLoop(9)
}

func printLoop(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d  ", i, j, i*j)
		}
		fmt.Println()
	}
}
