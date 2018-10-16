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
		if value, ok := element.(int); ok { //comma-ok 来判断element的原类型是不是int
			fmt.Printf("index:%d, value:%d element:%d\n", index, value, element)
			fmt.Println("value type", reflect.ValueOf(value).Type())
			fmt.Println("elemen type", reflect.ValueOf(element).Type())
		} else if value, ok := element.(string); ok {
			fmt.Printf("index:%d, value:%s element:%s\n", index, value, element)
			fmt.Println("value type", reflect.ValueOf(value).Type())
			fmt.Println("elemen type", reflect.ValueOf(element).Type())
		} else if value, ok := element.(Person); ok {
			fmt.Printf("index:%d, value:%s element:%s\n", index, value, element)
			fmt.Println("value type", reflect.ValueOf(value).Type())
			fmt.Println("elemen type", reflect.ValueOf(element).Type())
		} else {
			fmt.Printf("list[%d]different type", index)
		}
	}

}
