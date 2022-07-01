package parser

import (
	"crawler.com/oys/learngo/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(.*www\.zhenai\.com/zhenghun/[^"]+)"`)
)

func ParserCity(contents []byte) engine.ParserResult{
	matchsRes := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matchsRes {
		result.Items = append(result.Items, m[2])
		result.Requests = append(
			result.Requests,
			engine.Request{
					string(m[1]),
					engine.NilParser,
				},
			)
	}
	return result
}
