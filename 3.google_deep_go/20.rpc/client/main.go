package main

import (
	"fmt"
	"go_study/3.google_deep_go/20.rpc"
	"net"
	"net/rpc/jsonrpc"
)

/*
输出:
3.3333333333333335 <nil>
3.3333333333333335 division by zero
*/
func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)
	var result float64
	err = client.Call("DemonService.Div",
		rpcdemon.Args{10, 3}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemonService.Div",
		rpcdemon.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
