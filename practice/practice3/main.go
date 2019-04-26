package main

import (
	"fmt"
	"time"
)

const (
	Male = iota + 1
	Female
)

func main() {
	sec := time.Now().Unix()
	if sec%Female == 0 {
		fmt.Println(sec)
		fmt.Println("woman")
	} else {
		fmt.Println(sec)
		fmt.Println("man")
	}
}
