package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test0(b interface{}) {

	v, ok := b.(Student)
	fmt.Println(v, ok)
}


func test1(b interface{}) {
	//类型
	t := reflect.TypeOf(b)
	fmt.Println(t)

	//类别
	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		//%v:原始值
		//%T:类型
		fmt.Printf("%v %T\n", stu, stu)
	}
}


func testInt(b interface{}) {
	//val相当于1个指针
	val := reflect.ValueOf(b)
	//val.Elem() 相当于 *b  重点!!!!!!!!!!
	val.Elem().SetInt(100)
	c := val.Elem().Int()
	fmt.Printf("get value  interface{} %d\n", c)
	fmt.Printf("string val:%d\n", val.Elem().Int())
}

func main() {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92,
	}

	//{stu01 18 92} true
	test0(a)
	fmt.Println("====================")

	/*
	main.Student
	struct
	{stu01 18 92} main.Student
	*/
	test1(a)
	fmt.Println("====================")

	/*
	get value  interface{} 100
	string val:100
	100
	 */
	var b int = 1
	b = 200
	testInt(&b)//这里必须传地址
	fmt.Println(b)

}
