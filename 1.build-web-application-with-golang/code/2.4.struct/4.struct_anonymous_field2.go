package main

import (
	"fmt"
)

type Skills []string //数组切片

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	specialty string
}

func main() {
	//初始化学生jane
	jane := Student{Human: Human{"Jane", 35, 100},
		specialty: "Biology"}

	//现在我们来访问响应的字段
	fmt.Println("name is", jane.name)
	fmt.Println("age is", jane.age)
	fmt.Println("weight is", jane.weight)
	fmt.Println("specialty is", jane.specialty)

	jane.Skills = []string{"anatomy"}
	fmt.Println("skills ", jane.Skills)

	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("skills ", jane.Skills)

	jane.int = 123
	fmt.Println("number is ", jane.int)

}
