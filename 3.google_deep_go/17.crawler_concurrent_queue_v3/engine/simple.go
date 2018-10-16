package engine

import (
	"go_study/3.google_deep_go/17.crawler_concurrent_queue_v3/fetcher"
	"log"
	)


//简单版本engine
type SimpleEngine struct {}

func (e SimpleEngine)Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests)  > 0 {
		//dequeue
		r := requests[0]
		requests = requests[1:]

		//worker
		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		//enqueue
		requests = append(requests, parseResult.Requests...)

		//print item
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)// %v:意思是原来是什么就打印什么
		}
	}
}

func worker(r Request) (ParserResult, error) {
	//fetch
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url :%s %v",
			r.Url, err)
		return ParserResult{}, err
	}

	//parse
	return r.ParserFunc(body), nil
}