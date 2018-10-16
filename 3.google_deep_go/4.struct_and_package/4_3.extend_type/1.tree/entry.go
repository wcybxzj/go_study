package main

import (
	"fmt"
	"go_study/3.google_deep_go/4_3.extend_type/1.tree/tree"
)

//通过组合扩展原有类型，给tree加后续遍历
//实际就是把要扩充的struct包装在新的struct内
type myTreeNode struct {
	node *tree.Node
}

//receiver可传值也可传类型
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	//报错:
	//cannot call pointer method on myTreeNode literal
	//cannot take the address of myTreeNode literal
	//分析:
	// myTreeNode{myNode.node.Left} 叫myTreeNode literal
	//报错的原因是因为postOrder的receiver要求是指针，
	//而myTreeNode literal不行，需要一个变量
	//myTreeNode{myNode.node.Left}.postOrder()

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()

	myNode.node.Print()
}

func main() {
	//创建root
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //c语言root.right->left
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue3(4)

	//中序遍历
	root.Traverse()
	fmt.Println()

	//后续遍历
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
