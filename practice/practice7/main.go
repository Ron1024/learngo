package main

import (
	"fmt"
)

func main() {
	//for i := 101; i < 200; i++ {
	//	if sushu(i) {
	//		fmt.Println(i, " is a sushu !")
	//	}
	//}

	sum := 1
	for i := 1; i <= 3; i++ {
		sum *= i
	}
	fmt.Println(sum)

}

func sushu(a int) bool {
	//sqrt := int(math.Sqrt(float64(a)))
	//fmt.Println(sqrt)
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}

	return true
}

func shuixianhuashu(a int, b int) {

	for i := 100; i < 999; i++ {
		g := i / 100
		h := (i - g*100) / 10
		gh := i - g*100 - h*10

		if (g*g*g + h*h*h + gh*gh*gh) == i {
			fmt.Println("is a shuixianhuashu", i)
		}

	}

}

func jiesheng(a int) int {
	var sum int
	for i := 1; i < a; i++ {
		sum *= i
	}
	return sum
}
