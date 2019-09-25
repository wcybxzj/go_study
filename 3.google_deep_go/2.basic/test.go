package main

import "fmt"

func test1(outData map[string]interface{})  {
	outData["abc"]=13
}

func main() {
	dataOut := make(map[string]interface{})


	test1(dataOut)


	fmt.Println(dataOut)
}