package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println("This is in print() , ", node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value tu nil node . Ignored !")
		return
	}
	node.Value = value
	fmt.Println("This is in setValue , ", node)
}

