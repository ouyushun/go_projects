package parser

import (
	"crawler.com/oys/learngo/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParserResult{
	reg := regexp.MustCompile(cityListRe)
	matchsRes := reg.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matchsRes {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{string(m[1]), engine.NilParser, }, )
	}
	return result
}