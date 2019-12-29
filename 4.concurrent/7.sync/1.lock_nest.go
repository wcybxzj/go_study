package main

import (
	"sync"
)

type rpcSt struct {
	Host string
	rwLock sync.RWMutex
}

var rpc rpcSt


func InitGrpc() {
	rpc.rwLock.Lock()
	defer rpc.rwLock.Unlock()

	rpc.Host = ""
}


func rpcLogInfo(log string) () {
	rpc.rwLock.Lock()
	var hostEmpty bool
	hostEmpty = false
	if rpc.Host == ""{
		hostEmpty = true
	}
	rpc.rwLock.Unlock()

	if hostEmpty == true {
		InitGrpc()
	}
}

func LogDataGrpc() {
	go rpcLogInfo("111")
}


func main() {
	InitGrpc()
	rpcLogInfo("222")
}
