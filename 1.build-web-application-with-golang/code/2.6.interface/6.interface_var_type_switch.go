package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Element interface{}

type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return p.name + "---" + strconv.Itoa(p.age)
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		//element.(type) 语法不能在switch外的任何逻辑里面使用，
		//如果你要在switch外面判断一个类型就使用`comma-ok`。
		switch value := element.(type) {
		case int:
			fmt.Printf("index:%d, value:%d element:%d\n", index, value, element)
			fmt.Println("value type", reflect.ValueOf(value).Type())
			fmt.Println("element type", reflect.ValueOf(element).Type())
		case string:
			fmt.Printf("index:%d, value:%s element:%s\n", index, value, element)
			fmt.Println("value type", reflect.ValueOf(value).Type())
			fmt.Println("element type", reflect.ValueOf(element).Type())
		case Person:
			fmt.Printf("index:%d, value:%s element:%s\n", index, value, element)
			fmt.Println("value type", reflect.ValueOf(value).Type())
			fmt.Println("element type", reflect.ValueOf(element).Type())
		default:
			fmt.Printf("list[%d]different type", index)
		}
	}

}
