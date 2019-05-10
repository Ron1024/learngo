package main

func main() {
	root := &node{
		Name: "root",
	}

	left1 := &node{
		Name: "left1",
	}
	right1 := &node{
		Name: "Right1",
	}

	right11 := &node{
		Name: "right11",
	}
	left11 := &node{
		Name: "left11",
	}

	left111 := &node{
		Name: "left111",
	}

	root.left = left1
	root.right = right1

	left1.left = left11
	left1.right = right11
	left11.left = left111

	traverse(root)
}
