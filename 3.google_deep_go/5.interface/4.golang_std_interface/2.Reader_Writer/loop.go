package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//方法参数要求实现了io.Reader借口
//此时函数参数可以是文件名或者字符串
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//打印文件内容
	file, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	printFileContents(file)

	fmt.Println("=====================")

	//打印跨行字符串
	s := `abc "d"kk
	5555 `
	printFileContents(strings.NewReader(s))
}
