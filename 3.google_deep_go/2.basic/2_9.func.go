package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//测试1:函数
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		//使用err返回比panic更好程序不会停止
		return 0, fmt.Errorf("unsuported op:" + op)
	}
}

//测试2:
func div(a, b int) (q, r int) {
	//写法1:(推荐)
	//return a / b, a % b

	//写法2:
	q = a / b
	r = a % b
	return
}

//测试3:函数式编程重写eval,需要用到反射
//参数名:op 类型类型:func(int, int)int
func applay(op func(int, int) int, a, b int) int {
	//通过反射获得函数指针
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args"+
		"(%d, %d) result:", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//测试5:可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	//测试1:
	//正常情况
	fmt.Println(eval(13, 3, "/"))
	//错误情况
	if result, err := (eval(13, 3, "x")); err != nil {
		fmt.Printf("fail err:%s\n", err)
	} else {
		fmt.Printf("success result:%d\n", result)
	}
	fmt.Println("======================================")

	//测试2: div测试
	fmt.Println(div(13, 3))
	fmt.Println("======================================")

	//测试3:函数式编程
	//Calling function main.pow with args(3, 4) result:81
	//main的package name, pow是函数名
	fmt.Println(applay(pow, 3, 4))
	fmt.Println("======================================")

	//测试4:函数式编程-匿名函数
	//输出:Calling function main.main.func1 with args(3, 4) result:81
	//第一个main是package name,第二个main是函数名,第三个是匿名随机名
	fmt.Println(
		applay(
			func(a, b int) int {
				return int(math.Pow(float64(a), float64(b)))
			}, 3, 4))
	fmt.Println("======================================")

	//测试5:
	fmt.Println(sum(1, 2, 3, 4, 5))
}
