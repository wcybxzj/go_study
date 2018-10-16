package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

//测试1:
func test1() {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

//测试2:
//通过控制头部来实现访问wwww.imooc.com  302跳转到 m.imooc.com
func test2() {
	request, err := http.NewRequest(http.MethodGet,
		"http://www.imooc.com", nil) //get没有request body
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

//test3:
//fcuntion purpose:
// use http.Client.CheckRedirect to test 302

//CheckRedirect.via:
// redirect may happen many times,every redirect's path store in it

//CheckRedirect.req:
//every redirect target stroe in req
func test3() {
	request, err := http.NewRequest(http.MethodGet,
		"http://www.imooc.com", nil) //get没有request body
	request.Header.Add("User-Agent",
					"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request,
							via []*http.Request) error {
				fmt.Println("Redirect!!!!!!!!!!!!1:",req)
				return nil//use default redirect policy
		},
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}


func main() {
	//test1()
	test2()
}
