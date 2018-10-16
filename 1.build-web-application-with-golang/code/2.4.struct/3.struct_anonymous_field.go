package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      //匿名字段, 名字是隐匿的,但是类型还是需要的
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 25, 120}, "Compute Science"}
	//比c语言少写一个名字
	fmt.Println("name is", mark.name)
	fmt.Println("name is", mark.age)
	fmt.Println("name is", mark.weight)

	//更新数据
	mark.speciality = "AI"
	fmt.Printf("%s\n", mark.speciality)

	mark.weight++
	fmt.Printf("%d\n", mark.weight)
	fmt.Printf("%d\n", mark.Human.weight) //这样也可以
}
