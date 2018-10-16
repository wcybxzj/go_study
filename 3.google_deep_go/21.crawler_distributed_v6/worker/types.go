package worker

import (
	"errors"
	"fmt"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/config"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/engine"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/zhenai/parser"
	"log"
)

//原始的函数格式
//func ParseCity(contents []byte, _ string) engine.ParseResult {}  2个参数
//func ParserCityList(contents []byte, _ string) engine.ParseResult {} 2个参数
//func ParseProfile(contents []byte, url string, name string) engine.ParseResult {} 3个参数

//序列化后的函数json:
//因为ParseProfile是3参数,比其他函数出一个参数所以序列化的时候多一个userName
//{"ParseCity", nil},{"ParseCityList", nil},{"ParseProfile", userName}
type SerializedParser struct {
	Name string
	Args interface{}
}

//因为engine.Request里面是函数没法在jsonRPC中传递
//所以模拟它写了一个,方便在网络上传递
type Request struct {
	Url    string
	Parser SerializedParser
}

//同上
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

//序列化
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

//反序列化
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineRequest, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error DeserializeRequest request err:%v ", err)
			continue
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}

//将字符串转成函数
//方法1:将函数存入map, 然后通过名称获取函数
//方法2:switch case
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(
			parser.ParserCityList,
			config.ParseCityList), nil

	case config.ParseCity:
		return engine.NewFuncParser(
			parser.ParseCity,
			config.ParseCity), nil

	case config.NilParser:
		return engine.NilParser{}, nil

	case config.ParseProfile:
		if userName, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userName), nil
		} else {
			return nil, fmt.Errorf("invalid %v", p.Args)
		}
	default:
		re := fmt.Sprintf("unkown parser name:%v", p.Name)
		return nil, errors.New(re)
	}
}
