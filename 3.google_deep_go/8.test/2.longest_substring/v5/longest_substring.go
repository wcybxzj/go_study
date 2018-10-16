package main

import "fmt"

//将数组拿出来,应该是相当于静态数据,不是很理解
var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubstr(s string) int {
	//每次使用请要清理
	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI >= start {
			start = lastI
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}
	return maxLength
}

func test1() {
	//英文:成功
	fmt.Println(lengthOfNonRepeatingSubstr("abcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubstr("bbbb"))     //1
	fmt.Println(lengthOfNonRepeatingSubstr("pwwkew"))   //3
	fmt.Println(lengthOfNonRepeatingSubstr(""))         //0
	fmt.Println(lengthOfNonRepeatingSubstr("b"))        //1
	fmt.Println(lengthOfNonRepeatingSubstr("abcdef"))   //6
	//中文:成功
	fmt.Println(lengthOfNonRepeatingSubstr("这里是米克旺")) //6
	fmt.Println(lengthOfNonRepeatingSubstr("一二三二一"))  //3
}

func main() {
	test1()
}
