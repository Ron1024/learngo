package main

import "fmt"

func main() {
	s := []int{2, 7, 11, 15}
	var v []int
	v = twoSum(s, 9)
	fmt.Println(v)

}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if target-nums[i] == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nodes := ListNode{}
	current := &nodes
	overNum := 0
	for l1 != nil || l2 != nil {
		var l1Val int
		var l2Val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := overNum + l1Val + l2Val
		overNum = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if overNum > 0 {
		current.Next = &ListNode{Val: overNum}
	}
	return nodes.Next
}
