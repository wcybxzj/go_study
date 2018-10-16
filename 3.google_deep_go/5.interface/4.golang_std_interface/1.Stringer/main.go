package main

import (
	"fmt"
	"time"

	"go_study/3.google_deep_go/5.interface/4.golang_std_interface/1.String/mock"
	"go_study/3.google_deep_go/5.interface/4.golang_std_interface/1.String/real"
)

const url = "http://www.qq.com"

//借口
type Retriever interface {
	Get(url string) string //方法
}

//借口
type Poster interface {
	Post(url string, form map[string]string) string //方法
}

//组合借口
type RetrieverPoster interface {
	Retriever //借口
	Poster    //借口
	//Test(string)//定义其他的方法
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ybx",
			"course": "golang",
		})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

//type switch:查看接口类型和值
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T , Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

/*
输出:
Inspecting Retriever: {Contents=this is a fake imooc.com} //实现了String
> Type:*mock.Retriever , Value:Retriever: {Contents=this is a fake imooc.com} //实现了String
> Type switch: Contents: this is a fake imooc.com

Inspecting &{Mozila/5.0 1m0s}
> Type:*real.Retriever , Value:&{Mozila/5.0 1m0s}
> Type switch: UserAgent: Mozila/5.0

r is not a mock retriever
another faked imooc.com
*/
func main() {
	var r Retriever

	mockRetriever := mock.Retriever{
		Contents: "this is a fake imooc.com"}
	r = &mockRetriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozila/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	//Type assertion
	//此r是&real.Rertriever
	//类似c语言的assert,这里的意思是希望r借口变量是(*mock.Rertriever)
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println(session(&mockRetriever))
}
