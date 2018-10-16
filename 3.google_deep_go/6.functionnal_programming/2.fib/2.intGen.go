package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//想法:
//给函数实现接口io.Reader,让函数向文件那样能被循环打印
//做法:
//把fibonacci的返回函数做成一个类型,
//通过把函数定义为类型,
//就可以让类型去实现io.Reader接口
type intGen func() int

//给intGen来实现io.Reader接口的Read方法
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 { //设置上限,否则一直生成停止不了
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	//TODO:incorrect if p is too small
	return strings.NewReader(s).Read(p)
}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFile(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//让函数向文件那样可以被循环打印
	f := fibonacci()
	printFile(f)
	fmt.Println("================")
	//打印文件内容
	file, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	printFile(file)
	fmt.Println("================")
	//打印字符串
	s := `abc"def"
	kkkkk
	123

	p`
	fmt.Println(strings.NewReader(s))
	fmt.Println("================")
}
