package main

import (
	"fmt"
	"unicode"
)



func main() {
	s := "Hello 世界！"
	for _, r := range s {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c", r) // 世界
		}
	}
}
