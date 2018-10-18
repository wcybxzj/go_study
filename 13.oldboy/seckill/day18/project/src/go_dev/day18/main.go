package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
		var stu *Student
	var p People = stu
	if p == nil {
		fmt.Println("Yes")
	} else {
		fmt.Println("no")
	}
	result := live()
	if result == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
		fmt.Printf("%T %v\n", result, result)
	}
}