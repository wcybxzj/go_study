package main

import (
	"fmt"

	"go_study/3.google_deep_go/10.channel/4.tree/tree"
)

//使用channel来遍历tree
func main() {
	//创建root
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //c语言root.right->left
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue3(4)

	//打印最大的节点的数值
	//子协程去中序树并将树node值发送到channel
	//主协程从channel读取获取最大的node值
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Println("max value of node is :%d", maxNode)
}
