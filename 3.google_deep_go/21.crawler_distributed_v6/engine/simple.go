package engine

import (
	"log"
)

//简单版本engine
type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		//dequeue
		r := requests[0]
		requests = requests[1:]

		//worker
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		//enqueue
		requests = append(requests, parseResult.Requests...)

		//print item
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item) // %v:意思是原来是什么就打印什么
		}
	}
}
