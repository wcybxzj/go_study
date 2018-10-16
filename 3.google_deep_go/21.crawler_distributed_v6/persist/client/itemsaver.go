package client

import (
	"go_study/3.google_deep_go/21.crawler_distributed_v6/config"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/engine"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			//call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				//保存失败会报错
				log.Printf("item saver error"+
					"saving item %d: %v", itemCount, item)
			}
		}
	}()
	return out, nil
}
