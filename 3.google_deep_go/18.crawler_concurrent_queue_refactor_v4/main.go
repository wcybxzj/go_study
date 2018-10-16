package main

import (
	"go_study/3.google_deep_go/18.crawler_concurrent_queue_refactor_v4/engine"
	"go_study/3.google_deep_go/18.crawler_concurrent_queue_refactor_v4/zhenai/parser"
	"go_study/3.google_deep_go/18.crawler_concurrent_queue_refactor_v4/scheduler"
)

//第4版并发爬虫：
//知识点1:重构第三版并发爬虫+队列调度器

//知识点2:获取上海的页内容
func main() {
	//并发爬虫第2版本+并发调度器
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:10,//worker数量
	}

	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.QueuedScheduler{},
	//	WorkerCount:10,//worker数量
	//}

	//默认城市列表页
	//e.Run(
	//	engine.Request{
	//		Url:"http://city.zhenai.com",
	//		ParserFunc:parser.ParserCityList,
	//})

	//只获取上海用户列表
	e.Run(engine.Request{
		Url: "http://city.zhenai.com/shanghai",
		ParserFunc:parser.ParseCity,
	})

}