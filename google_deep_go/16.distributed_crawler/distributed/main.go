package main

import (
	"net/rpc"
	"strings"

	"log"

	"flag"

	"go_study/google_deep_go/16.distributed_crawler/crawler/config"
	"go_study/google_deep_go/16.distributed_crawler/crawler/engine"
	"go_study/google_deep_go/16.distributed_crawler/crawler/scheduler"
	"go_study/google_deep_go/16.distributed_crawler/crawler/zhenai/parser"
	itemsaver "go_study/google_deep_go/16.distributed_crawler/distributed/persist/client"
	"go_study/google_deep_go/16.distributed_crawler/distributed/rpcsupport"
	worker "go_study/google_deep_go/16.distributed_crawler/distributed/worker/client"
)

// 添加启动配置的参数
var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String(
		"worker_hosts", "",
		"worker hosts (comma separated)")
)

// 分布式版
func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	// todo:字符串的输入需要做校验，但暂时就不做了
	// 建立client连接池
	//pool := createClientPool(strings.Split(*workerHosts, ","))
	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("Error connecting to %s: %v", h, err)
		}
	}
	// 创建完之后，进行分发
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
