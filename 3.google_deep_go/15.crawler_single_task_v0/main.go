package main

import (
	"regexp"
	"fmt"
	"go_study/3.google_deep_go/15.crawler_single_task_v0/engine"
	"go_study/3.google_deep_go/15.crawler_single_task_v0/zhenai/parser"
)

//单任务爬虫
//只抓取每个城市,第一个列表页中,用户详细页信息
func main() {
	engine.Run(
		engine.Request{
			Url:"http://city.zhenai.com",
			ParserFunc:parser.ParserCityList,
	})
}

//target html: <a href="http://city.zhenai.com/zaozhuang" class="">枣庄</a>
//[a-z0-9]+: means is get zaozhuang
//[^>]: means not >
//*: 0 or many
//[^>]+: means 0 or many not > charset
//[^<]+: means 1 or many not < charset
func test1(contents []byte)  {
	re := regexp.MustCompile(`<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	//example1:
	//FindAll return [][]byte
	//[]byte: mean is string
	//[][]byte:mean is string arr
	matches := re.FindAll(contents, -1)
	for _, m := range matches{
		fmt.Printf("%s\n",m)
	}
	fmt.Printf("matches found:%d\n",len(matches))
}
