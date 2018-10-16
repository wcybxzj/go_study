package main

import "fmt"

type person struct {
	name string
	age  int
}

func test1() {
	var p person
	p.name = "ybx"
	p.age = 25
	fmt.Printf("%s:%d\n", p.name, p.age)
}

func test2() {
	p := person{"Tom", 25}
	fmt.Printf("%s:%d\n", p.name, p.age)
	p = person{age: 31, name: "man"}
	fmt.Printf("%s:%d\n", p.name, p.age)
}

//指针类型
func test3() {
	p := new(person)
	*p = person{age: 18, name: "mm"}
	fmt.Printf("%s:%d\n", p.name, p.age)
}

func main() {
	test1()
	test2()
	test3()
}
