package main

import (
	"fmt"
)

func reverse1(str string) string {
	var result string
	strLen := len(str)

	for i := strLen - 1; i >= 0; i-- {
		result = result + fmt.Sprintf("%c", str[i])
	}

	return result
}

func reverse2(str string) string {
	var result []byte
	tmp := []byte(str) //string->byte slice
	length := len(str)
	for i := length - 1; i >= 0; i-- {
		result = append(result, tmp[i])
	}
	return string(result)
}

func main() {
	v1 := reverse1("abcdef")
	fmt.Println(v1)

	v2 := reverse2(v1)
	fmt.Println(v2)
}
