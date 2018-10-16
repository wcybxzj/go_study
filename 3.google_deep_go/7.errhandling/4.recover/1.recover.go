package main

import (
	"fmt"
)

//recover的返回值是interface{}
//所以需要做type assertions判断
//r这个借口变量是不是error类型
func tryRecover() {
	defer func() {
		r := recover()

		if r != nil {
			if err, ok := r.(error); ok {
				//r确实是error类型
				fmt.Println("(((((错误发生)))))):", err)
			} else {
				//不是预期的类型再次panic(r)
				//程序也不知道如何处理了
				panic(fmt.Sprintf("I dont know what to do:%v", r))
			}
		}
	}()

	//测试1:
	//编译错误直接不能编译

	//a := 5 / 0

	//测试2:
	//运行时错误,可以用recover来恢复
	//(((((错误发生)))))): runtime error: integer divide by zero
	//必须写在这里上面的recover才能处理

	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	//测试3
	//自定义错误
	//(((((错误发生)))))): This is  an error
	//必须写在这里上面的recover才能处理

	//panic(errors.New("This is  an error"))

	//测试4:
	//recover不知道如何处理,再次抛出错误
	//必须写在这里上面的recover才能处理
	//输出:
	//panic: 123 [recovered]
	//	panic: I dont know what to do:123

	//panic(123)

	//测试5:
	//没有错误
	//fmt.Println("1111111")

	fmt.Println("main finish!(只要出现了panic程序就不会往panic下面走了)")
}

func main() {
	tryRecover()
}
