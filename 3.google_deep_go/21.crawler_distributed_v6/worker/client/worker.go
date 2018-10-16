package client

import (
	"go_study/3.google_deep_go/21.crawler_distributed_v6/config"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/engine"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		//请求序列化
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		//RPC调用
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		//结果反序列化
		return worker.DeserializeResult(sResult), nil
	}
}
