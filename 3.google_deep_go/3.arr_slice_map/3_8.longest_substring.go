package main

import "fmt"

//例:寻找最长的不含有重复字符的字串
//abcabcbb --> abc
//bbbbbbb  --> b
//pwwkew   -->wke

/*
格式:
lastOccurred[字符]=字母最后出现的位置
start:当前string的识别位置的最初位置
*/

/*
算法:老师的算法很好只是没精力去看,我自己的版本是C语言那个版本
对于每个字符x:
if lastOccurred[x]不存在 OR lastOccurred[x] < start{
	//不用操作
}else if lastOccurred[x] >= start{
	//更新start 到 lastOccurred[x]+1的位置
}
更新lastOccurred[x],更新maxLength
*/
//第一版:仅支持ascii
func lengthOfNonRepeatingSubstrV1(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	//[]byte(s):string强转成[]byte (在处理多字节语言时候有问题,后边会讲rune)
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

//第二版:国际化支持多语言
func lengthOfNonRepeatingSubstrV2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	//[]byte(s):string强转成[]byte (在处理多字节语言时候有问题,后边会讲rune)
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func test1() {
	//英文:成功
	fmt.Println(lengthOfNonRepeatingSubstrV1("abcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubstrV1("bbbb"))     //1
	fmt.Println(lengthOfNonRepeatingSubstrV1("pwwkew"))   //3
	fmt.Println(lengthOfNonRepeatingSubstrV1(""))         //0
	fmt.Println(lengthOfNonRepeatingSubstrV1("b"))        //1
	fmt.Println(lengthOfNonRepeatingSubstrV1("abcdef"))   //6
	//中文:失败
	fmt.Println(lengthOfNonRepeatingSubstrV1("这里是米克旺")) //15
	fmt.Println(lengthOfNonRepeatingSubstrV1("一二三二一"))  //5

	fmt.Println("===============================================")
	//英文:成功
	fmt.Println(lengthOfNonRepeatingSubstrV2("abcabcbb")) //3
	fmt.Println(lengthOfNonRepeatingSubstrV2("bbbb"))     //1
	fmt.Println(lengthOfNonRepeatingSubstrV2("pwwkew"))   //3
	fmt.Println(lengthOfNonRepeatingSubstrV2(""))         //0
	fmt.Println(lengthOfNonRepeatingSubstrV2("b"))        //1
	fmt.Println(lengthOfNonRepeatingSubstrV2("abcdef"))   //6
	//中文:成功
	fmt.Println(lengthOfNonRepeatingSubstrV2("这里是米克旺")) //6
	fmt.Println(lengthOfNonRepeatingSubstrV2("一二三二一"))  //3

}

func main() {
	test1()
}
