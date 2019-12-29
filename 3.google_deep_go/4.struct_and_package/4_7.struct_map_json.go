package main

import (
	"encoding/json"
	"fmt"
)

//map->json
//json->map
func test1(){
	m := map[string]interface{} {"name":"taoge", "age":30, "addr":"China"}
	fmt.Println(m) //打印map: map[addr:China age:30 name:taoge]

	data, _ := json.Marshal(m)//map -> []byte
	fmt.Println(string(data)) //[]byte -> string: {"addr":"China","age":30,"name":"taoge"}

	m1 := make(map[string]interface{})
	_ = json.Unmarshal(data, &m1) // json_string -> map
	fmt.Println(m1) //map[addr:China age:30 name:taoge]
}

//struct->json
//json->struct
func test2(){
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Addr string `json:"addr"`
	}

	p1 := Person{
		Name: "taoge",
		Age:  30,
		Addr: "China",  // oh my god, this comma cannot be omitted
	}

	fmt.Println(p1)//打印结构体: {taoge 30 China}

	data, _ := json.Marshal(p1) //struct -> []byte
	fmt.Println(string(data))// {"name":"taoge","age":30,"addr":"China"}

	var p2 Person
	_ = json.Unmarshal(data, &p2)
	fmt.Println(p2)
}

//Marshal():可以将struct者map 转成 []byte
//string():可以将[]byte转换成string
//Unmarshal():可以将[]byte 转成 struct或者map
func main() {
	test1()
	fmt.Println("-------------")
	test2()
}