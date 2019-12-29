package main

import (
	"errors"
	"fmt"
)

//只是返回错误非空
func test1() error {
	var val error = errors.New("XXX")
	return val

}

func main() {
	var err error
	for  {
		err = test1()
		if err != nil {
			break
		}
	}
	fmt.Println(err.Error())
}
