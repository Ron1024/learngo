package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string

	fmt.Scanf("%s\n", &str)

	testCoverToInt(str)
}

func testCoverToInt(str string) {
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("This string can not be convert")
	} else {
		fmt.Println(number)
	}
}
