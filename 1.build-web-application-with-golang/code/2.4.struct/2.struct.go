package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person
	tom.name, tom.age = "Tom", 18

	bob := person{age: 25, name: "Bob"}

	paul := person{"Paul", 43}

	tb_older, tb_diff := Older(tom, bob)
	tp_older, tp_diff := Older(tom, paul)
	bp_older, bp_diff := Older(bob, paul)

	fmt.Printf("%s compare %s ,older  is %s  beyond %d years\n",
		tom.name, bob.name, tb_older.name, tb_diff)
	fmt.Printf("%s compare %s ,older  is %s  beyond %d years\n",
		tom.name, paul.name, tp_older.name, tp_diff)
	fmt.Printf("%s compare %s ,older  is %s  beyond %d years\n",
		bob.name, paul.name, bp_older.name, bp_diff)
}
