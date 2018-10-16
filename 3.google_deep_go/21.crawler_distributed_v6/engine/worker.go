package engine

import (
	"go_study/3.google_deep_go/21.crawler_distributed_v6/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {

	//fetch
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher error fetching url :%s %v",
			r.Url, err)
		return ParseResult{}, err
	}

	//parse
	//return r.ParserFunc(body, r.Url), nil
	return r.Parser.Parse(body, r.Url), nil
}
