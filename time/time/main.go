package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getCurrentTime())
}

func getCurrentTime() string {
	return time.Now().Format("2017/06/15 08:05:00")
}
