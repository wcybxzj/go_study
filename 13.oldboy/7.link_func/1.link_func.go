package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) say() person {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	return p
}

func main() {
	p := &person{"ybx", 123}
	p.say().say().say().say().say().say()
}
