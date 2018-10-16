package main

import (
	"fmt"
	"reflect"
)

/*
type: float64
kind is float64: true
value: 3.4
*/
func test1() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

//func test2() {
//	var x float64 = 3.4
//	v := reflect.ValueOf(x)
//	v.SetFloat(7.1) //不能这样修改
//}

func test3() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)
	fmt.Println("now xi si", x)
}

func main() {
	test1()
	test3()

}
