package engine

import (
	"crawler.com/oys/learngo/fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	fmt.Println(seeds)
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(req.Url)

		if err != nil {
			log.Printf("Fetcher error url : %s, error: %s", req.Url, err)
			continue
		}

		parserResult := req.ParserFunc(body)
		//fmt.Printf("%s", requests)
		requests = append(requests, parserResult.Requests...)
		//fmt.Printf("%s", requests)
		//fmt.Printf("Got item %v", parserResult.Items)
	}
}
