package main

import (
	"fmt"

	"log"

	"flag"

	"go_study/google_deep_go/16.distributed_crawler/distributed/rpcsupport"
	"go_study/google_deep_go/16.distributed_crawler/distributed/worker"
)

// 创建启动配置参数,第三个参数就是我们的解析参数解析语句,使用go run xxxx.go --help即可查询
var port = flag.Int("port", 0, "The port for me to Listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must Specify a Port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
