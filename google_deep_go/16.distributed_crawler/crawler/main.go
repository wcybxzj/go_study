package main

import (
	"go_study/google_deep_go/16.distributed_crawler/crawler/config"
	"go_study/google_deep_go/16.distributed_crawler/crawler/engine"
	"go_study/google_deep_go/16.distributed_crawler/crawler/persist"
	"go_study/google_deep_go/16.distributed_crawler/crawler/scheduler"
	"go_study/google_deep_go/16.distributed_crawler/crawler/zhenai/parser"
)

// 单机版
func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
