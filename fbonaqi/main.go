package main

import "fmt"

func main() {
	n := fbnq(1000)
	fmt.Println(n)
}

func fbnq(num int) []int64 {
	var a []int64
	a = make([]int64, num)
	a[0] = 1
	a[1] = 1

	for i := 2; i < num; i++ {
		a[i] = int64((i - 1) + (i - 2))
	}
	return a
}
