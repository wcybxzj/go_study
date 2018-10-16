package main

import (
	"go_study/google_deep_go/15.single_and_concurrent_crawler/crawler/config"
	"go_study/google_deep_go/15.single_and_concurrent_crawler/crawler/engine"
	"go_study/google_deep_go/15.single_and_concurrent_crawler/crawler/persist"
	"go_study/google_deep_go/15.single_and_concurrent_crawler/crawler/scheduler"
	"go_study/google_deep_go/15.single_and_concurrent_crawler/crawler/zhenai/parser"
)

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
		Url: "http://city.zhenai.com",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
