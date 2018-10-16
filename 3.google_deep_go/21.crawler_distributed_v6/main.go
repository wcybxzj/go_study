package main

import (
	"flag"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/config"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/engine"
	itemsaver "go_study/3.google_deep_go/21.crawler_distributed_v6/persist/client"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/rpcsupport"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/scheduler"
	worker "go_study/3.google_deep_go/21.crawler_distributed_v6/worker/client"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/zhenai/parser"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)

	if err != nil {
		panic(err) //這裏需要用panic,如果itemSaver
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      10, //worker数量
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	//默认城市列表页
	e.Run(
		engine.Request{
			Url: "http://city.zhenai.com",
			Parser: engine.NewFuncParser(
				parser.ParserCityList,
				config.ParseCityList),
		})
}

//clients是私有数据,是无法共享给其他协程的
//通过channel来进行把私有数据进行共享
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
