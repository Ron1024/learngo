package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Getenv("GOPATH")

	sysinfo := os.Getenv("GOARCH")
	fmt.Println("GOPATH : ", path)
	fmt.Println("GOARCH : ", sysinfo)
}
