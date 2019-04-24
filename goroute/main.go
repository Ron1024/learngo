package main

import (
	"fmt"
	"learngo/goroute/calc"
)

func main() {
	//for i := 0; i < 100; i++ {
	//	go print2.PrintCh(i)
	//}

	var pipe = make(chan int)
	go calc.Add(100, 200, pipe)
	go calc.Add(200, 300, pipe)
	for j := 0; j < len(pipe); j++ {
		sum := <-pipe
		fmt.Println(sum)
	}

	//time.Sleep(10 * time.Second)

}
