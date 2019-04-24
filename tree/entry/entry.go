package main

import (
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right :=myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {

	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}

	root.Right.Left = new(tree.Node)

	//node := []tree.Node{
	//	{Value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(root)
	//fmt.Println(node)
	//root.Print()
	//root.Left.Print()
	//root.Right.Left.Print()
	//
	//root.SetValue(5)
	//
	//root.Print()
	//
	//fmt.Println("This is out of setValue() , ", root)
	//
	//var pRoot *tree.Node
	//pRoot.SetValue(200)
	//pRoot = &root
	//pRoot.SetValue(300)
	//pRoot.Print()

	root.Traverse()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}
