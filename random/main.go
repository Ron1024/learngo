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

	//生成10个随机整数
	for i := 0; i < 10; i++ {
		fmt.Println("10个随机整数:", rand.Int())
	}
	//生成在一定范围内的整数
	for i := 0; i < 10; i++ {
		fmt.Println("范围内的10个随机整数:", rand.Intn(10))
	}

	for i := 0; i < 10; i++ {
		fmt.Println("10个随机浮点数(32):", rand.Float32())
	}

	for i := 0; i < 10; i++ {
		fmt.Println("10个随机浮点数(64):", rand.Float64())
	}

}
