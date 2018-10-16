package parser

import (
	"regexp"
	"go_study/3.google_deep_go/18.crawler_concurrent_queue_refactor_v4/engine"
	)

//<td><a href="http://album.zhenai.com/u/1607404583" target="_blank">南情哥哥</a></td>
var (
	profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	//<a class="next-page" href="http://city.zhenai.com/shanghai/2">下一页</a>
	//或者
	//相亲栏目:<a target="_blank" href="http://city.zhenai.com/shanghai/gongwuyuan">上海公务员相亲</a>
	cityUrlRe = regexp.MustCompile(`href="(http://city.zhenai.com/shanghai/[^"]+)"`)
)

//处理每个城市的用户列表页
func ParseCity(contents []byte) engine.ParserResult {
	//log.Printf("begin of ParseCity");

	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	//第二次感受到闭包的巧妙:

	//问题:ParseProfile是两个参数,
	//原先engine.Requst.ParserFunc定义的是1个参数,
	//ParseProfile不能作用在可原先engine.Requst.ParserFunc

	//写法1:不使用闭包
	//修改点1:修改engine.Requst.ParserFunc成2个参数,
	//修改点2:并且还要修改 ParserCityList() ParseCity()改成2参数,即使没有也要定义,还要改调用的地方

	//写法2:使用闭包
	//原来结构体engine.Requst.ParserFunc不用改
	//原来相关函数(ParserCityList() ParseCity())都不用改
	//闭包注意点:闭包中如果使用循环中的变量要复制给别的变量

	for _, m := range matches{
		url := m[1]
		name := m[2]//必须怎么写否则闭包中的m[2]将执行同一个名字
		//result.Items = append(result.Items, "User "+string(name))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url: string(url),
				ParserFunc: func(contents [] byte) engine.ParserResult {
					return ParseProfile(contents, string(name))//不能使用m[2]否则一直都是循环最后的名字
				},
			})
	}

	matches = cityUrlRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:string(m[1]),
				ParserFunc:ParseCity,
			})
	}

	//log.Printf("end of ParseCity");
	return result
}