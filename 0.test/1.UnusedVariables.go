package main

import "fmt"

var gvar int //not an error

func main() {
	var one int //error, unused variable
	one = 123
	fmt.Println(one)

	two := 2 //error, unused variable
	fmt.Println(two)

	//var three int //error, even though it's assigned 3 on the next line
	three = 3
}
