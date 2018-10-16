package engine

import (
	"go_study/3.google_deep_go/15.crawler_single_task_v0/fetcher"
	"log"
	)

func Run(seeds ...Request)  {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests)  > 0 {
		//dequeue
		r := requests[0]
		requests = requests[1:]

		//fetch
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher error fetching url :%s %v",
				r.Url, err)
			continue
		}

		//parse
		parseResult := r.ParserFunc(body)

		//enqueue
		requests = append(requests, parseResult.Requests...)

		//print item
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)// %v:意思是原来是什么就打印什么
		}
	}
}
