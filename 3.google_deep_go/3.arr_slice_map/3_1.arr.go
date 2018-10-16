package main

import "fmt"

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println("==================")
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	fmt.Println("==================")
	for index, value := range arr3 {
		fmt.Printf("index:%d, value:%d\n", index, value)
	}
	fmt.Println("==================")
	for index := range arr3 {
		fmt.Printf("index:%d\n", index)
	}

	fmt.Println("==================")
	//二维数组4行5列
	var grid [4][5]int
	fmt.Println(grid)
}
