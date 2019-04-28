package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	num := rand.Intn(100)
	testCheck(num)
}

func testCheck(n int) {
	for {
		var input int
		fmt.Println("请输入：")
		fmt.Scanf("%d\n", &input)
		flag := false
		switch {
		case input < n:
			fmt.Println("小了，重新输入")
		case input == n:
			fmt.Println("你猜对了")
			flag = true
		case input > n:
			fmt.Println("大了，重新输入")
		}
		if flag {
			break
		}
	}
}
