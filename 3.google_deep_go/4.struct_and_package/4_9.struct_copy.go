package main

import (
	"fmt"
)

type Dog struct {
	age  int
	name string
}

func test1()  {
	roger := Dog{5, "Roger"}
	mydog := roger
	if roger == mydog {
		fmt.Println("Roger and mydog are equal structs")
	}

	roger.age=100
	roger.name="name123"
	fmt.Println(mydog.age)
}

func test2_1(DogPointer *Dog)  {
	DogPointer.age=123
}

func test2_2(Dog Dog)  {
	Dog.age=456
}

func test2()  {
	roger := Dog{5, "Roger"}

	test2_1(&roger)
	fmt.Println(roger.age)

	test2_2(roger)
	fmt.Println(roger.age)
}

func main() {
	//test1()
	test2()
}