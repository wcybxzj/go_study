package main

import "fmt"

func test1(cols ...string){
	fmt.Println(cols)
}


func test2(arr []string){
	fmt.Println(arr)
}

func main() {
	//test1("aaa","bbb")

	var s []string
	s = append(s, "aaaaaa")
	s = append(s, "bb")

	test2(s)

}
