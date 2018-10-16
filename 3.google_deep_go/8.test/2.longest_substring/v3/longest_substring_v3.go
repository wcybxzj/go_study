package main

import "fmt"

//第三版:性能优化V2版本用slice替换map
//但是benmark性能更差
func lengthOfNonRepeatingSubstrV3(s string) int {
	lastOccurred := make([]int, 0xffff) //优化成数组,长度65536,大量垃圾回收更慢了
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
	fmt.Println(lengthOfNonRepeatingSubstrV3("abcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubstrV3("bbbb"))     //1
	fmt.Println(lengthOfNonRepeatingSubstrV3("pwwkew"))   //3
	fmt.Println(lengthOfNonRepeatingSubstrV3(""))         //0
	fmt.Println(lengthOfNonRepeatingSubstrV3("b"))        //1
	fmt.Println(lengthOfNonRepeatingSubstrV3("abcdef"))   //6
	//中文:成功
	fmt.Println(lengthOfNonRepeatingSubstrV3("这里是米克旺")) //6
	fmt.Println(lengthOfNonRepeatingSubstrV3("一二三二一"))  //3
}

func main() {
	test1()
}
