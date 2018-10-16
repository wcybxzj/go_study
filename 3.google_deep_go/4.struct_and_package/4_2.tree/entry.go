package main

import (
	"fmt"

	"go_study/3.google_deep_go/4.struct_and_package/4_2.tree/tree"
)

func main() {
	//创建root
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) //c语言root.right->left

	//创建nodes
	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	//打印root
	root.Print()

	//receiver传值或者传地址
	//可以看到无论setValueX()的调用者是指针或者值都可以调用函数
	root.Right.Left.SetValue1(4) //值不能修改
	root.Right.Left.Print()      //0
	root.Right.Left.SetValue2(4) //值能修改
	root.Right.Left.Print()      //4

	//打印nodes
	fmt.Println(nodes)

	//打印root
	root.Print() //3
	root.SetValue2(100)

	//root传给一个指针
	pRoot := &root
	pRoot.Print() //100

	pRoot.SetValue2(200)
	pRoot.Print() //200

	//nil指针
	var pRoot2 *tree.Node
	pRoot2.SetValue3(300)
	pRoot2 = &root
	pRoot2.SetValue3(400)

	pRoot2.Print()

	fmt.Println("==========================")
	root.Traverse()
}
