package client

import (
	"net/rpc"

	"go_study/google_deep_go/16.distributed_crawler/crawler/engine"
	cconfig "go_study/google_deep_go/16.distributed_crawler/distributed/config"
	"go_study/google_deep_go/16.distributed_crawler/distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	//client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	//if err != nil {
	//	return nil, err
	//}
	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(cconfig.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
