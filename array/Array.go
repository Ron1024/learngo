package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)
	printArray(arr1)
	testTowDimensionalArray()
}

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
func testTowDimensionalArray() {
	arr := [][]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}}
	for row, val := range arr {
		for col, vals := range val {
			fmt.Printf("(%d,%d)=%d ", row, col, vals)
		}
		fmt.Println()
	}
}
