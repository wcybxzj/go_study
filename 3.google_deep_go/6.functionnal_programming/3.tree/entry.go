package main

import (
	"fmt"

	"go_study/3.google_deep_go/6.functionnal_programming/3.tree/tree"
)

func main() {
	//创建root
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //c语言root.right->left
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue3(4)

	root.Traverse()

	nodeCount := 0

	root.TraverseFunc(
		func(node *tree.Node) {
			nodeCount++
		})
	fmt.Println("Node count:", nodeCount)
}
