package main

import "fmt"

func main() {
	var b int
	Test(b)
}

func Test(a interface{}) {
	b := a.(int)
	b += 3
	fmt.Println(b)
}
