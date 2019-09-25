package main

import (
	"fmt"
	"time"
	"go_study/3.google_deep_go/5.interface/1.interface/mock"
	"go_study/3.google_deep_go/5.interface/1.interface/real"
)

//获取借口,借口中都是方法
type Retriever interface {
	Get(url string) string
}

//根据借口的规范使用各种类型的实现
func download(r Retriever) string {
	return r.Get("http://www.qq.com")
}

//测试两个Retriever
func test1() {
	var r Retriever
	r = mock.Retriever{"God is me"}
	fmt.Println(download(r))

	//实现时候指定了receiver是指针 必须用指针调用
	r = &real.Retriever{}
	fmt.Println(download(r))
}

//查看借口的类型和内容:
/*
输出:
----------------------------
类型:mock.Retriever 值:{God is me}
Contents: God is me
----------------------------
类型:*real.Retriever 值:&{Mozilla/5.0 1m0s}
UserAgent: Mozilla/5.0
----------------------------
1m0s
----------------------------
not a mockRetriever

*/
func test2() {
	//方法1:
	var r Retriever
	//对象值拷贝到r
	r = mock.Retriever{"God is me"}
	fmt.Println("----------------------------")
	type_switch(r)

	//对象的指针拷贝到r
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Println("----------------------------")
	type_switch(r)

	//方法2: Type assertion
	fmt.Println("----------------------------")
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	fmt.Println("----------------------------")
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mockRetriever")
	}
}

//方法1:type switch
func type_switch(r Retriever) {
	fmt.Printf("类型:%T 值:%v\n", r, r) //T:是看类型，v:类型内容
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}

func main() {
	test1()
	fmt.Println("==============================")
	test2()
}
