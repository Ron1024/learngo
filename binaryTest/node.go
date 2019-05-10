package main

import "fmt"

type node struct {
	Name        string
	left, right *node
}

func traverse(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node)
	traverse(node.left)
	traverse(node.right)
}
