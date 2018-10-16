package main

import (
	"go_study/13.oldboy/day9/chat/chat_server/lib"
	"time"
)

func main() {
	lib.InitRedis("localhost:6379", 16, 1024, time.Second*300)
	lib.InitUserMgr()
	lib.RunServer("0.0.0.0:10000")
}
