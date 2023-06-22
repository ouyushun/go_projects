package engine

import (
	"crawler.com/oys/learngo/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct {

}

func (eg SimpleEngine) Run(seeds ...Request) {
	fmt.Println(seeds)
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		parserResult, err := worker(req)
		if err == nil {
			log.Printf("%v", err)
			continue
		}
		requests = append(requests, parserResult.Requests...)
	}
}


func worker(req Request) (ParserResult, error){
	//fmt.Println("fetching ", req.Url)
	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		log.Printf("Fetcher error url : %s, error: %s", req.Url, err)
		return ParserResult{}, err
	}
	parserResult := req.ParserFunc(body)
	return parserResult, nil
}