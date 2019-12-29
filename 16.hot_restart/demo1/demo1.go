package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
kill -INT 1687
*/


//3种解决 当子协程接受到 INT信号,让父进程等待子进程结束才执行

//办法1:很不科学 等多长时间都不合适
func func1()  {
	for j:=0; j<10; j++ {
		fmt.Println("j:",j)
		time.Sleep(time.Second*10)
	}
}

//办法2:科学
func func2 (c chan struct{} )  {
	<-c
}


//https://developer20.com/golang-tips-and-trics-iii/
func main() {
	fmt.Println("demo1 pid:",os.Getpid())
	var srv http.Server
	srv.Addr=":12345"

	//这个channel很有用
	//当kill -INT pid发来的时候
	//可以让main协程等待子协程执行完返回
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		fmt.Println("优雅关闭")

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}

		for i:=0; i<3; i++ {
			fmt.Println("child,i:", i)
			time.Sleep(time.Second)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	fmt.Println("main block")
	func1()
	//func2(idleConnsClosed)
	fmt.Println("main unblock")
}
