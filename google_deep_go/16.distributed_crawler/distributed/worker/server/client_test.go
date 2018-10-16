package main

import (
	"testing"
	"time"

	"fmt"

	cconfig "go_study/google_deep_go/16.distributed_crawler/crawler/config"
	"go_study/google_deep_go/16.distributed_crawler/distributed/config"
	"go_study/google_deep_go/16.distributed_crawler/distributed/rpcsupport"
	"go_study/google_deep_go/16.distributed_crawler/distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/108906739",
		Parser: worker.SerializedParser{
			Name: cconfig.ParseProfile,
			Args: "安静的雪",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
