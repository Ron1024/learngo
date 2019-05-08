package main

import (
	"fmt"
	"sort"
)

func main() {
	testSortInt()
	testSortString()
	testSortFloat()

	testSearch()
}

//测试整数排序
func testSortInt() {
	var a = [...]int{1, 4, 2, 3, 5, 111, 5, 7, 12, 34}
	sort.Ints(a[:])
	fmt.Println(a)
}

func testSortString() {
	var str = [...]string{"acd", "abc", "我"}
	sort.Strings(str[:])
	fmt.Println(str)
}

func testSortFloat() {
	var str = [...]float64{3.4444, 2.1111, 4.2222, 1.1111, 0.1, 11.0}
	sort.Float64s(str[:])
	fmt.Println(str)
}

//测试查找
func testSearch() {
	var a = [...]int{1, 4, 2, 3, 5, 111, 5, 7, 12, 34}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 5)
	fmt.Println(a)
	fmt.Println(index)
}
