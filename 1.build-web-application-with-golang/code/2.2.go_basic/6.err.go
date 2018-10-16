package main

import "fmt"
import "errors"

func test1() {
	err := errors.New("errrrrrr")
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	test1()
}
