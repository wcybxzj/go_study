package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func testReadString() {
	file, err := os.OpenFile("/tmp/test.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	//write
	writer := bufio.NewWriter(file)
	len, err := writer.WriteString("123456")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success write len:", len)
	}
	writer.Flush()

	//read
	file.Seek(0, os.SEEK_SET)
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed, err:", err)
		return
	}

	fmt.Printf("read str succ, ret:%s\n", str)
}

//使用readline读取一行超级长的数据存在变量中
func testReadLine() {
	file, err := os.OpenFile("/tmp/bigline.log", os.O_CREATE|os.O_RDWR|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	//1.创建超级长的一行,写入文件
	writer := bufio.NewWriter(file)
	for i := 0; i < 2000; i++ {
		len, err := writer.WriteString("12345678abcdefghijklmnopqrstuvwxyz")
		if err != nil {
			panic(err)
		} else {
			fmt.Println("success write len:", len)
		}
	}
	len, err := writer.WriteString("ybx")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success write len:", len)
	}
	writer.Flush()

	//2.从文件读取超长的一行
	file.Seek(0, os.SEEK_SET)
	reader := bufio.NewReader(file)

	for {
		//isFinishOneLine 一行是否读取完毕
		str, isFinishOneLine, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("read string failed, err:", err)
				return
			}
		}

		fmt.Println("isFinishOneLine", isFinishOneLine)
		fmt.Printf("read str succ, ret:%s\n", str)
	}
}

func main() {
	//testReadString()
	testReadLine()
}
