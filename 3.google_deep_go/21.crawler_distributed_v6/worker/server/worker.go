package main

import (
	"flag"
	"fmt"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/rpcsupport"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

/*
帮助:
go run worker.go --help

启动3个woker服务:
go run worker.go --port=4001
go run worker.go --port=4002
go run worker.go --port=4003
*/
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(
		rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", *port),
			worker.CrawlService{}),
	)
}
