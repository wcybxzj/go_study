package main

import "fmt"

func test1()  {
	var s1 [] string
	s1 = append(s1, "aaa")
	s1 = append(s1, "bbb")
	fmt.Println(s1)
}

//output:
//[1 3 4 5 5]
//[1 3 4 5]
func test2() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := a1

	a1 = append(a1[:1], a1[2:]...)

	fmt.Println(a2)
	fmt.Println(a1)
}

func test3(s []int) ([]int){
	s1 := append(s[:1], s[2:]...)
	return s1
}

func test4()  {
	pureExcelData := make([][]string, 0)


	fmt.Println(len(pureExcelData))
}


func main() {
	//test1()
	//test2()


/*	a1 := []int{1, 2, 3, 4, 5}
	a2 := test3(a1)
	fmt.Println(a2)*/

	test4()


}
