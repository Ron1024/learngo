package main

import "fmt"

func main() {
	num := 100
	testPointer(&num)
	fmt.Println(num)
}
func testPointer(num *int) {
	fmt.Println(&num)
	*num = 100
}
