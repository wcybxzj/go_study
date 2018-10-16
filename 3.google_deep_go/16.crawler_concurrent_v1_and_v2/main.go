package main

import (
	"go_study/3.google_deep_go/16.crawler_concurrent_v1_and_v2/engine"
	"go_study/3.google_deep_go/16.crawler_concurrent_v1_and_v2/zhenai/parser"
	"go_study/3.google_deep_go/16.crawler_concurrent_v1_and_v2/scheduler"
)

//版本：
//并发爬虫第一版本:简单调度器(失败)
//并发爬虫第二版本:并发调度器(成功)
//第一版和第二版的核心区别就是scheduler/simple.go中Submit 提交请求给 in channel是否 每个request开一个协程

//数据:
//只抓取470个城市+每个城市只取得第一页 20个用户+20个用户详情页的信息
//470x20=8140个用户+8140个用户详情=16000个item
func main() {
	//使用简单串行版本engine:
	//engine.SimpleEngine{}.Run(
	//	engine.Request{
	//		Url:"http://city.zhenai.com",
	//		ParserFunc:parser.ParserCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:10,//worker数量
	}

	e.Run(
		engine.Request{
			Url:"http://city.zhenai.com",
			ParserFunc:parser.ParserCityList,
	})
}