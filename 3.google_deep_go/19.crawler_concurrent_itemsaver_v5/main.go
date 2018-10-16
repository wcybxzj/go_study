package main

import (
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/engine"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/persist"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/scheduler"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err) //這裏需要用panic,如果itemSaver
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10, //worker数量
		ItemChan:    itemChan,
	}

	//默认城市列表页
	e.Run(
		engine.Request{
			Url:        "http://city.zhenai.com",
			ParserFunc: parser.ParserCityList,
		})
}
