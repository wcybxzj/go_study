package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1 string
	s2 string
	s3 string
}

func (n NotknownType) String() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}


/*
main.NotknownType
struct
Field 0: Ada
Field 1: Go
Field 2: Oberon
[Ada-Go-Oberon]
*/
//反射:一个不知道什么类型的借口,通过反射机制来知道他是什么类型,什么值,里面有什么字段和方法
func main() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	fmt.Println(typ)

	knd := value.Kind() // struct
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//value.Field(i).SetString("C#")
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}