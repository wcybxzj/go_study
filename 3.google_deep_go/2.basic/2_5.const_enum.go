package main

import (
	"fmt"
)

func enums() {
	const (
		cpp        = iota //0
		_                 //1
		python            //2
		golang            //3
		javascript        //4
	)

	//b,kb, mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	enums()
}
