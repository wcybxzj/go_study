package main

import (
	"fmt"
	"time"
)

//dead lock
func test0() {
	c := make(chan int)
	v := <-c
	fmt.Println(v)

	//相当于pause()
	for {
		time.Sleep(100e9) //100秒
	}
}

//dead lock
func test1() {
	c := make(chan int)

	select {
	case v := <-c:
		fmt.Println(v)
	}

	//相当于pause()
	for {
		time.Sleep(100e9) //100秒
	}
}

//ok
func test2() {
	c := make(chan int)

	go func() {
		v := <-c
		fmt.Println(v)
	}()

	//相当于pause()
	for {
		time.Sleep(100e9) //100秒
	}
}

//ok
func test3() {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			}
		}
	}()

	//相当于pause()
	for {
		time.Sleep(100e9) //100秒
	}
}

//ok
func test4() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(3 * time.Second):
				fmt.Println("select timeout 3 second")
				o <- true
				break
			}
		}
	}()

	<-o
	fmt.Println("程序结束")
}

//ok
func test5() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)
			case <-time.After(3 * time.Second):
				fmt.Println("select timeout 3 second")
				o <- true
				break
			}
		}
	}()

	for i := 0; i < 5; i++ {
		c <- i
		time.Sleep(time.Second)
	}

	<-o
	fmt.Println("程序结束")
}

func main() {
	//test0()
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}
