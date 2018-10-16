package main

import (
	"fmt"
	"os"
)

func osTest1()  {
	fmt.Fprintf(os.Stdout, "write!!")
}

func osTest2()  {
	file, err := os.OpenFile("/tmp/123.txt", os.O_CREATE|os.O_WRONLY, 0664)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Fprint(file, "111111111111111111\n")
}

func main() {
	osTest1()
	osTest2()
}