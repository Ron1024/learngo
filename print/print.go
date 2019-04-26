package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1000
	s := "200000"
	//转换二进制
	fmt.Printf("%b\n", a)
	//打印十进制
	fmt.Printf("%d\n", a)
	//打印八进制
	fmt.Printf("%o\n", a)
	//小写的十六进制
	fmt.Printf("%x\n", a)
	//大写的十六进制
	fmt.Printf("%X\n", a)
	//带0x的十六进制
	fmt.Printf("%#x\n", a)
	//6位小数点
	fmt.Printf("%f\n", float32(a))
	//字符串
	fmt.Printf("%s\n", strconv.Itoa(a))
	//字符串转整形
	if b, err := strconv.ParseInt(s, 8, 32); err != nil {
		panic(err)
	} else {
		fmt.Printf("%d", b)
	}

	//fmt.Println(bits.ReverseBytes64(1000))

}
