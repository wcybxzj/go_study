package parser

import (
	"go_study/3.google_deep_go/15.crawler_single_task_v0/engine"
	"regexp"
	)

//	<a href="http://city.zhenai.com/aba" class="">阿坝</a>
const cityListRe  = `<a href="(http://city.zhenai.com/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//处理城市列表页
func ParserCityList(contents []byte) engine.ParserResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	//只取得1个城市,因为有470个城市,
	//单任务如果都取出来，半天也看不到最后的用户详情页
	limit :=10

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

		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
