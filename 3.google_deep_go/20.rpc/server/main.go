package main

import (
	"go_study/3.google_deep_go/20.rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*手动调试方法:
telnet localhost 1234
输入:
{"method":"DemonService.Div","params":[{"A":3, "B":0}], "id":123}
输出:
{"id":1,"result":0.75,"error":null}

输入:
{"method":"DemonService.Div","params":[{"A":3, "B":0}], "id":123}
输出:
{"id":123,"result":null,"error":"division by zero"}
*/
//简单的RPC服务器
func main() {
	rpc.Register(rpcdemon.DemonService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v", err)
			continue
		}

		//执行任务,就是rpc.go中的Div函数
		go jsonrpc.ServeConn(conn)
	}

}
