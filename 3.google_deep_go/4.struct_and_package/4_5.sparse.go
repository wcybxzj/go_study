package main

import (
	"fmt"

	"golang.org/x/tools/container/intsets"
)

func testSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	fmt.Println(s.Has(1))
	fmt.Println(s.Has(1000))
}

func main() {
	testSparse()
}
