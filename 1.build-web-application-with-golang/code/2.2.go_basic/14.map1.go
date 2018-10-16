package main

import "fmt"

//1.map的定义
//可以使用内建函数 make 也可以使用 map 关键字来定义 Map:
//如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
func test1() {
	//定义方法1:
	//var map1 map[string]string
	//map1 = make(map[string]string)

	//定义方法2:
	map1 := make(map[string]string)

	map1["Frecnch"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India"] = "新德里"

	for country := range map1 {
		fmt.Println(country, "首都是", map1[country])
	}

	var capital string
	capital, ok := map1["美国"]
	if ok {
		fmt.Println("美国首都是", capital)
	} else {
		fmt.Println("美国首都没找到")
	}
}

//delete删除map的一个元素
func test2() {
	map1 := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"China":  "Beijing"}

	fmt.Println("原始地图")

	for country := range map1 {
		fmt.Println(country, "首都是", map1[country])
	}

	fmt.Println("\n")
	delete(map1, "France")

	for key := range map1 {
		fmt.Println(key, "首都是:", map1[key])
	}
}

//map和slice一样也是引用类型,
//如果两个map同时指向一个底层,1个改了另外一个也会变
func test3() {
	m := make(map[string]string)
	m["Hello"] = "ybx"
	m1 := m
	m1["Hello"] = "wc"
	fmt.Println(m["Hello"])  //wc
	fmt.Println(m1["Hello"]) //wc
}

func main() {
	//test1()
	//test2()
	test3()
}
