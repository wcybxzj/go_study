package engine

import (
	"log"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/fetcher"
)

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
	return r.ParserFunc(body, r.Url), nil
}