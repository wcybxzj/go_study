package main

import "fmt"

func main() {
	var str = "hello world\n"

	//原封不动的输出
	var str1 = `
	床前明月光，
	疑是地上霜。
	举头望明月，
	我是郭德纲。`

	var b byte = 'c'

	fmt.Println(str) // hello world
	fmt.Println(str1)
	fmt.Println(b)        // 99
	fmt.Printf("%c\n", b) // c
}
