package test

import "fmt"

var (
	Name string = "ron"
	Age  int    = 10
)

func init() {
	fmt.Println("This the test package init().")
	fmt.Println("This is the test package init(),name = ", Name)
	fmt.Println("This is the test package init(),age = ", Age)

	Age = 1000
	fmt.Println("This is the test package init(),Updated the age ! age = ", Age)
}
