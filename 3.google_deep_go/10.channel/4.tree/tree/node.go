package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node Node) SetValue1(value int) {
	node.Value = value
}

func (node *Node) SetValue2(value int) {
	node.Value = value
}

func (node *Node) SetValue3(value int) {
	if node == nil {
		fmt.Println("setting value to nil " + "node Ignored!!!!!!!!!!!!!!!")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
