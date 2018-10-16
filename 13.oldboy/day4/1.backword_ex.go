package main

import (
	"fmt"
	)

//只支持英文
func process_english(str string) bool {
	for i:=0; i<len(str); i++{
		if i==len(str)/2 {
			break
		}
		last :=len(str)-i-1

		if str[i]!=str[last] {
			return false
		}
	}

	return true
}

//rune支持中文
func process_rune(str string)  {
	t := []rune(str)
	for i, v := range t{
		fmt.Printf("%v %v %d\n", i, v, len(string(v)))
	}
}

//只支持中文
func process_chinese(str string) bool {
	t := []rune(str)
	length := len(t)

	for i,_ := range t{
		if i==length/2 {
			break
		}
		last :=length-i-1

		if t[i]!=t[last] {
			return false
		}
	}
	return true
}

//测试1:中文失败
func test1()  {
	re := process_english("abc")//false
	fmt.Println(re)
	re = process_english("abcba")//true
	fmt.Println(re)
	re = process_english("中国中")//false 应该是true
	fmt.Println(re)
}

//测试2:rune
func test2()  {
	process_rune("中1国2中")
}

//测试3:中文成功
func test3()  {
	re := process_chinese("abc")//false
	fmt.Println(re)
	re = process_chinese("abcba")//true
	fmt.Println(re)
	re = process_chinese("中国中")//true
	fmt.Println(re)
	re = process_chinese("中1国1中")//true
	fmt.Println(re)
	re = process_chinese("中国人")//false
	fmt.Println(re)
}

func main() {
	//test1()
	//test2()
	test3()
}
