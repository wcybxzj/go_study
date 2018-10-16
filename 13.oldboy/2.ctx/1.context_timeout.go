package __ctx

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

/*
测试1:访问一个不能访问的网站引起超时
输出:Timeout! err: Get http://google.com: net/http: request canceled while waiting for connection
执行顺序:main协程发现超时, Done能读取到然后通知子协程结束
*/
func test1() {
	//创建上下文,并且设置2秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//取消上下文
	defer cancel()
	tr := &http.Transport{}
	client := http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		fmt.Println("http request failed, err", err)
		return
	}

	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()

	select {
	//超时
	case <-ctx.Done():
		//取消http请求
		tr.CancelRequest(req)
		res := <-c
		fmt.Println("Timeout! err:", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, err := ioutil.ReadAll(res.r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server Response:%s", out)
	}
}

/*
测试2:
i: 0
i: 1
Timeout!

执行顺序:main协程发现超时, Done能读取到然后通知子协程结束
*/
func test2() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //创建上下文,并且设置2秒超时
	defer cancel()                                                          //取消上下文

	c := make(chan int, 1)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("i:", i)
			time.Sleep(time.Second)
		}
		c <- 123
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Timeout! ")
	case res := <-c:
		fmt.Printf("Server Response:%v", res)
	}

	//time.Sleep(time.Second * 10)
}

//context超时例子
func main() {
	//test1()
	test2()
}
