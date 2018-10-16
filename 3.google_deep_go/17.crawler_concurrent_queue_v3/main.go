package main

import (
	"go_study/3.google_deep_go/17.crawler_concurrent_queue_v3/engine"
	"go_study/3.google_deep_go/17.crawler_concurrent_queue_v3/zhenai/parser"
	"go_study/3.google_deep_go/17.crawler_concurrent_queue_v3/scheduler"
)


func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:10,//worker数量
	}

	e.Run(
		engine.Request{
			Url:"http://city.zhenai.com",
			ParserFunc:parser.ParserCityList,
	})
}