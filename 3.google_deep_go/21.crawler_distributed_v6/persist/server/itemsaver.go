package main

import (
	"flag"
	"fmt"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/config"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/persist"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

//go run itemsaver.go --port=4000
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	//docker 运行必须用SetSniff(false)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
