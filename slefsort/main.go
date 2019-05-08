package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 9, 5, 3, 2, 1, 6, 4}
	//brr := selectionSort(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func bubbleSort(a []int) []int {
	if len(a) == 0 {
		return a
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				//var temp int
				//temp = a[i]
				//a[i] = a[j+1]
				//a[j+1] = temp
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func selectionSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		minIndex := i //设置最小值下标，依次向右

		//从最右侧的所有元素中选择最小的
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	val := arr[left]
	k := left

	for i := left + 1; i <= right; i++ {
		if arr[i] < val {
			arr[k] = arr[i]
			arr[i] = arr[k+1]
			k++
		}

	}
	fmt.Println(arr)
	arr[k] = val
	quickSort(arr, left, k-1)
	quickSort(arr, k+1, right)
}
