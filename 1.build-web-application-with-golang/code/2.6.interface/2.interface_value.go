package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//传说中的面向接口编程
//只要类型中扩展了这两个函数就可以
type Men interface {
	SayHi()
	Sing(word string)
}

func main() {
	mike := Student{Human{"mike", 25, "19810110101"}, "MIT", 0.00}
	paul := Student{Human{"paul", 26, "110"}, "Harvard", 100}
	sam := Employee{Human{"sam", 36, "120"}, "Golang", 1000}
	tom := Employee{Human{"tom", 37, "13034380999"}, "Things", 5000}

	var i Men

	//i能存储Student
	i = mike
	i.SayHi()
	i.Sing("mike is me !")

	//i能存储Employee
	i = tom
	i.SayHi()
	i.Sing("Born to be wild")

	//定义slice Men
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

}
