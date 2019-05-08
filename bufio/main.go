package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("error to readline!")
	}
	//wordCount, spaceCount, numberCount, otherCount := process(string(res))
	//
	//fmt.Println("worldCount:", wordCount)
	//fmt.Println("spaceCount:", spaceCount)
	//fmt.Println("numberCount:", numberCount)
	//fmt.Println("otherCount:", otherCount)
	str := strings.Split(string(res), "+")
	result := mult(strings.TrimSpace(str[0]), strings.TrimSpace(str[1]))

	fmt.Println(result)
}

func process(str string) (worldCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, val := range t {
		switch {
		case val >= 'a' && val <= 'z':
			fallthrough
		case val >= 'A' && val <= 'Z':
			worldCount++
		case val == ' ':
			spaceCount++
		case val >= 0 && val <= 0:
			numberCount++
		default:
			otherCount++

		}
	}
	return
}

//大数相加实现
//实现方法，对于字符串进行逆序相加，判断是否有进位，如果进位加到下一位中
func mult(str1, str2 string) (result string) {
	index1 := len(str1) - 1
	index2 := len(str2) - 1
	fmt.Println(str1, str2)
	var left int //进位
	//先处理低位数据处理到低位最小位数结束
	for index1 >= 0 && index2 >= 0 {
		num1 := str1[index1] - '0'
		num2 := str2[index2] - '0'
		sum := int(num1) + int(num2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}

		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		num1 := str1[index1] - '0'
		sum := int(num1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}
	for index2 >= 0 {
		num2 := str2[index2] - '0'
		sum := int(num2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}
	if left != 0 {
		result = fmt.Sprintf("1%s", result)
	}

	return result
}
