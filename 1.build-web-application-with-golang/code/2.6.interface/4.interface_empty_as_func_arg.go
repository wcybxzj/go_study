package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

//只要任何类型实现了String(),就可用于fmt.Println()
func (h Human) String() string {
	return h.name + "-" + strconv.Itoa(h.age) + "-" + h.phone
}

//注：实现了error接口的对象（即实现了Error() string的对象），fmt输出时，会调Error()，因此不必再定义String()
func main() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("human is", Bob)
}
