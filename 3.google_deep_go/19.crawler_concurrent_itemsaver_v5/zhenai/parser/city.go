package parser

import (
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/engine"
	"regexp"
)

//<td><a href="http://album.zhenai.com/u/1607404583" target="_blank">南情哥哥</a></td>
var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	//<a class="next-page" href="http://city.zhenai.com/shanghai/2">下一页</a>
	//或者
	//相亲栏目:<a target="_blank" href="http://city.zhenai.com/shanghai/gongwuyuan">上海公务员相亲</a>
	cityUrlRe = regexp.MustCompile(`href="(http://city.zhenai.com/shanghai/[^"]+)"`)
)

//处理每个城市的用户列表页
func ParseCity(contents []byte, _ string) engine.ParserResult {
	//log.Printf("begin of ParseCity");

	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matches {
		//result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ProfileParser(string(m[2])),
			})
	}

	matches = cityUrlRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	//log.Printf("end of ParseCity");
	return result
}
