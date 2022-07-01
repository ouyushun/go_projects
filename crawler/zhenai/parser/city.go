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
		name := string(m[2])
		result.Items = append(result.Items, m[2])
		result.Requests = append(
			result.Requests,
			engine.Request{
					string(m[1]),
				func(contents []byte) engine.ParserResult {
					return ParseProfile(contents, name)
				},
				},
			)
	}
	return result
}
