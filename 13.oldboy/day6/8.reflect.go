package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"student_name"`
	Age   int
	Score float32
	Sex   string
}

func (s Student) Print() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end----")
}

func (s Student) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	//类型必须是指针
	//指针必须执行struct
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取结构体字段
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("stu1000")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}
	fmt.Printf("struct has %d fields\n", num)

	//获取结构体的tag,利用这个自己也可以实现struct->json
	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	//结构体方法数量
	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//调用结构体方法
	var params []reflect.Value
	val.Elem().Method(0).Call(params)
}


func test1(){
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}

	TestStruct(&a)
	fmt.Println(a)
}

func test2()  {
	var a Student = Student{
		Name:  "stu01",
		Age:   18,
		Score: 92.8,
	}

	//struct -> json:使用的也是放射机制
	result, _ := json.Marshal(a)
	fmt.Println("json result:", string(result))
}

func main() {
	test1()
	test2()
}