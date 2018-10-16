package parser

import (
	"go_study/3.google_deep_go/16.crawler_concurrent_v1_and_v2/engine"
	"regexp"
	)

//	<a href="http://city.zhenai.com/aba" class="">阿坝</a>
const cityListRe  = `<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//处理城市列表页
func ParserCityList(contents []byte) engine.ParserResult{
	//log.Printf("ParserCityList");
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matches{
		url := m[1]
		city := m[2]
		result.Items = append(result.Items, "City "+string(city))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(url),
				ParserFunc: ParseCity,
			})
	}

	return result
}
