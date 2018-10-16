package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
	fmt.Println()
}

//知识点5:值接受者vs指针接受者
//要改变内容必须使用指针接受者
//如果参数内容过大要考虑是否要用指针代替值做接受者

//指针接受者
//receiver 如果像这里进行值传递，是不能修改node里的value的
func (node Node) SetValue1(value int) {
	node.Value = value
}

//值接受者
func (node *Node) SetValue2(value int) {
	node.Value = value
}

//知识点6:
//nil指针也可以调用方法
func (node *Node) SetValue3(value int) {
	if node == nil {
		fmt.Println("setting value to nil " + "node Ignored!!!!!!!!!!!!!!!")
		return
	}
	node.Value = value
}

//知识点1:
//go仅支持封装, 不支持继承和多态

//知识点2:
//go没有构造函数,要自己写工厂函数

//知识点3:
//c/c++就不能怎么写,因为Node是函数的局部变量,返回地址就会报错

//知识点4:
//go语言可以返回局部变量的地址
//如果数据不需要return 地址,Node就会在stack上分配
//如果数据return 地址,Node就会在heap上分配
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
