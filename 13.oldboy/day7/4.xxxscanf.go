package main

import (
	"fmt"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                = "%f / %d / %s"
)

func test1()  {
	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

func test2()  {
	var str = "stu01 18 89.92"
	var stu student
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}

func test3()  {
	fmt.Printf("Hi %s %s!\n", "dabing", "loveyou")
}

func main() {
	//test1()
	//test2()
	test3()
}
