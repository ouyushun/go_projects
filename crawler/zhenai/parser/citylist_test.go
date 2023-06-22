package parser

import (
	"crawler.com/oys/learngo/fetcher"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := fetcher.Fetch("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		return
	}
	ParserCityList(contents)
}