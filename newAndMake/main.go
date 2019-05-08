package main

import "fmt"

func main() {
	var i int

	fmt.Println(i)
	j := new(int)
	fmt.Println(*j)
	k := make(chan int, 1)
	fmt.Println(k)

}
