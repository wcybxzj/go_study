package main

import (
	"fmt"

	"go_study/3.google_deep_go/5.interface/3.interface_combination/mock"
)

const url = "http://www.qq.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
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

//组合借口
type RetrieverPoster interface {
	Retriever
	Poster
	//Test(string)//定义其他的方法
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func test() {
	retriever := mock.Retriever{"this is a fake imooc.com"}
	fmt.Println("try a session")
	fmt.Println(session(&retriever))
}

func main() {
	test()
}
