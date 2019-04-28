package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := testHasPreffix("http://www.baidu.com", "http://")
	//
	//fmt.Println("str", str)
	//
	//str2 := testHasSuffix("http://www.baidu.com", ".com")
	//
	//fmt.Println("str2", str2)
	//
	//str3 := testStrngsIndex("http://www.baidu.com", ".com")
	//
	//fmt.Println("str3", str3)

	testStringReplace("http://www.baidu.com", ".com")

	fmt.Println(testStringCount("http://www.baidu.com", "w"))

	fmt.Println(testStringRepeat("你好", 4))

	fmt.Println(testStringToLowerCase("WOSHI"))
	fmt.Println(testStringToUpperCase("woshi"))
	fmt.Println(testStringsTrimSpace(" w o s h i    "))
	fmt.Println(testStringsLeftTrimSpace(" w o s h i    "))
	fmt.Println(testStringsRightTrimSpace("    w o s h i    "))
	fmt.Println(testStringsField("w o s h i"))
	fmt.Println(testStringsSplit("w,o,s,h,i", ","))
	fmt.Println(testStringsJoin(testStringsField("w o s h i"), "-"))

}

func testHasPreffix(s string, s2 string) string {
	var stria string
	if strings.HasPrefix(s, s2) {
		fmt.Println(s, "has preffix ", s2)
	} else {
		stria = s2 + s
	}
	return stria
}

func testHasSuffix(s string, s2 string) string {
	if strings.HasSuffix(s, s2) {
		fmt.Println(s, "has sufferfix ", s2)
	} else {
		s = s + s2
	}
	return s
}

func testStrngsIndex(s string, str string) string {
	num := strings.Index(s, str)
	return string(s[num])
}

func testStringReplace(str string, str1 string) {
	str = strings.Replace(str, "www", "WWW", 1)
	fmt.Println(str)
}

func testStringCount(str string, substr string) int {
	return strings.Count(str, substr)
}
func testStringRepeat(str string, count int) string {
	return strings.Repeat(str, count)
}
func testStringToUpperCase(str string) string {
	return strings.ToUpper(str)
}

func testStringToLowerCase(str string) string {
	return strings.ToLower(str)
}
func testStringsTrimSpace(str string) string {
	return strings.TrimSpace(str)
}
func testStringsLeftTrimSpace(str string) string {
	return strings.TrimLeft(str, " ")
}
func testStringsRightTrimSpace(str string) string {
	return strings.TrimRight(str, " ")
}

func testStringsField(str string) []string {
	return strings.Fields(str)
}
func testStringsSplit(str string, split string) []string {
	return strings.Split(str, split)
}
func testStringsJoin(str []string, sep string) string {
	return strings.Join(str, sep)
}
