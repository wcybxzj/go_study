package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//测试1:十进制转二进制
func convertToBinary(num int) string {
	result := ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		panic(fmt.Sprintf("wrong num:%d", num))
	}

	for ; num > 0; num /= 2 {
		last_binary := num % 2
		result = strconv.Itoa(last_binary) + result
	}
	return result
}

//测试2:循环读取文件
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //相当于while(xxx)
		fmt.Println(scanner.Text())
	}
}

//测试3:死循环
func forever() {
	for { //很简单,因为go经常要用死循环
		fmt.Println("abc")
		time.Sleep(1e9) //1 second
	}
}

func main() {

	fmt.Println(
		convertToBinary(5),  //101
		convertToBinary(13), //1101
	)
	fmt.Println("================")

	printFile("abc.txt")

	fmt.Println("================")
	forever()

}
