package main

import (
	"fmt"
	"time"
)

type St struct{
	age int
}
//send nil to a chnnel
/*
输出:
Read:<nil>
Read:<nil>
...............
*/
func oneTestSendNil()  {
	ch := make(chan *St)

	go func(c chan *St) {
		for {
			select {
			case v:=<-c:
				fmt.Printf("Read:%v\n",v)
			}
		}
	}(ch)


	go func(c chan *St) {
		for i:=0; i<10000; i++{
			c<-nil
			time.Sleep(1*time.Second)
		}
	}(ch)

	time.Sleep(100*time.Second)
}

func main() {
	oneTestSendNil()
}