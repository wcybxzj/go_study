package main

import "fmt"

//测试1:
//map的两种格式
//格式1:普通map
//map[K]V
//格式2:复合型map,就map1里值是map2
//map1[K1]map2[K2]V
func test1() {
	//方法1:
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "immoc",
		"quality": "notbad",
	}
	fmt.Println(m)
	fmt.Println("=====================================")
	//方法2:
	m2 := make(map[string]int) //m2 == empty map
	fmt.Println(m2)
	fmt.Println("=====================================")
	//方法3:
	var m3 map[string]int //m3 == nil
	fmt.Println(m3)
	fmt.Println("=====================================")
	//遍历:
	//每次运行输出的顺序不一样,原因是map是无序的hashtable
	//如果想要保证顺序需要手动对key进行排序，存到一个slice然后进行排序
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("=====================================")
	//获取特定值,值存在
	courseName := m["course"]
	fmt.Println(courseName)
	fmt.Println("=====================================")
	//获取特定值,值不存在,输出一个空行,不会报错
	val := m["not_exist"]
	fmt.Println(val) //golang变量会自动初始化成zeroValue
	fmt.Println("=====================================")
	//判断一个key是否存在
	val1, ok := m["course"]
	fmt.Println(val1, ok) //golang true

	val2, ok := m["not_exist"]
	fmt.Println(val2, ok) //空 false

	if courseName, ok := m["course"]; ok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key is not exist")
	}
	fmt.Println("=====================================")
	//删除元素
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	m["new"]="11111"
	fmt.Println(m)
}

//map的key
//map使用哈希做底层，key需要可以比较
//除了slice,map,function的其他内建类型都可以做为key
//struct类型不包含上述字段,也可以作为key

func main() {
	test1()
}
