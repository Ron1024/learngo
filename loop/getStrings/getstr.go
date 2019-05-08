package main

import "fmt"

func main() {

	str := "Hello ,中国！"
	for length, val := range str {
		fmt.Printf("index[%d] val[%c] len[%d] \n", length, val, len([]byte(string(val))))
	}
}
