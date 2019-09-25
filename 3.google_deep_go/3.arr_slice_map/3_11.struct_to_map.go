package main

import (
	"fmt"
	"reflect"
)


type Student  struct {
	Id int `json:"id"`
	Name string `json:"age_size"`
}


func StructToMapDemo(obj interface{}) map[string]interface{}{
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}

func TestStructToMap() {
	student := Student{10, "jqw"}
	data := StructToMapDemo(student)
	fmt.Println(data)
}

func main() {
	TestStructToMap()
}