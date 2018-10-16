package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ybx123@gmail.com"

const text2 = "My email is ybx123@gmail.com@shou.com"

const text3 = `My email is ybx123@gmail.com@shou.com
email1 is abc@def.org
email2 is kk@qq.com
email3 is ddd@abc.com.cn
`

/*
输出:
ybx123@gmail.com
*/
func test1() {
	//测试1:
	//输入正则表达式,获取正则表达式对象
	//re, err := regexp.Compile("ybx123@gmail.com")

	//因为输入正则表达式是用户输入,一定可以获取正则对象
	re := regexp.MustCompile("ybx123@gmail.com")
	match := re.FindString(text)
	fmt.Println(match)
}

/*
输出:
My email is ybx123@gmail.com.cn
My email is ybx123@gmail.com.cn
*/
func test2() {
	//转义字符:\n
	//正则中点:是任意内容
	//要表示普通的点要用\.

	//方法1:
	//""会对其中内容进行转义,所以"\."这个转义字符不存在就会报错
	//"\\"意思是说明\是个普通的斜杠
	//最终"\\."表达的意思:普通的\和普通的点
	re1 := regexp.MustCompile(".+@.+\\..+")
	match1 := re1.FindString(text)
	fmt.Println(match1)

	//方法2:
	//使用``不对内容进行转义
	re2 := regexp.MustCompile(`.+@.+\..+`)
	match2 := re2.FindString(text)
	fmt.Println(match2)
}

/*
输出:
ybx123@gmail.com
==================================
ybx123@gmail.com@shou.com
==================================
ybx123@gmail.com
==================================
*/
func test3() {
	re2 := regexp.MustCompile(`[a-zA-Z0-9]+@.+\..+`)
	match2 := re2.FindString(text)
	fmt.Println(match2)
	fmt.Println("==================================")

	match2 = re2.FindString(text2)
	fmt.Println(match2)
	fmt.Println("==================================")

	re2 = regexp.MustCompile(
		`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match2 = re2.FindString(text2)
	fmt.Println(match2)
	fmt.Println("==================================")
}

//测试4:匹配多行内容
/*
输出:
ybx123@gmail.com
==================================
[ybx123@gmail.com abc@def.org kk@qq.com ddd@abc.com]
==================================
[ybx123@gmail.com abc@def.org kk@qq.com ddd@abc.com.cn]
==================================
*/
func test4() {
	re := regexp.MustCompile(
		`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text3)
	fmt.Println(match)
	fmt.Println("==================================")

	match2 := re.FindAllString(text3, -1)
	fmt.Println(match2)
	fmt.Println("==================================")
	//第二个[]中的\. 可以直接写成. 
	re3 := regexp.MustCompile(
		`[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+`)

	match3 := re3.FindAllString(text3, -1)
	fmt.Println(match3)
	fmt.Println("==================================")
}

//搜索子匹配 返回的是一个二维的数组
/*
output:

[ybx123@gmail.com ybx123 gmail com]
[abc@def.org abc def org]
[kk@qq.com kk qq com]
[ddd@abc.com.cn ddd abc.com cn] //notice here
==================================
[ybx123@gmail.com ybx123 gmail com]
[abc@def.org abc def org]
[kk@qq.com kk qq com]
[ddd@abc.com.cn ddd abc com.cn]//notice here
==================================
[ybx123@gmail.com ybx123 gmail .com]
[abc@def.org abc def .org]
[kk@qq.com kk qq .com]
[ddd@abc.com.cn ddd abc .com.cn] //notic here
==================================
*/
func test5 ()  {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text3, -1)
	for _, m := range match{
		fmt.Println(m)
	}
	fmt.Println("==================================")

	re2 := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	match2 := re2.FindAllStringSubmatch(text3, -1)
	for _, m := range match2{
		fmt.Println(m)
	}
	fmt.Println("==================================")

	re3 := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match3 := re3.FindAllStringSubmatch(text3, -1)
	for _, m := range match3{
		fmt.Println(m)
	}
	fmt.Println("==================================")
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}
