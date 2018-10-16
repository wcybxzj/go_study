package main

import (
	"fmt"
)

type treeNode struct {
	value int
}

//知识点1:
//值针接受者,可接受值/指针
//指针接受者,可接受值/指针

//知识点2:
//只有用指针做接受者的才可以修改数据内容
func (node treeNode) func1(val int) {
	node.value = val
}

func (node *treeNode) func2(val int) {
	node.value = val
}

/*
输出:
go run 4_1.value_and_pointer_receiver.go
{3} &{5} &{0}
{3} &{5} &{0}
{111} &{222} &{333}
*/
func main() {
	var node1 treeNode
	node1 = treeNode{value: 3} //值
	node2 := &treeNode{5}      //指针
	node3 := new(treeNode)     //指针
	fmt.Println(node1, node2, node3)

	node1.func1(123)
	node2.func1(456)
	node3.func1(789)
	fmt.Println(node1, node2, node3)

	node1.func2(111)
	node2.func2(222)
	node3.func2(333)
	fmt.Println(node1, node2, node3)
}
