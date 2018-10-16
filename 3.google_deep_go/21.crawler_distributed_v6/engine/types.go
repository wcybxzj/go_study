package engine

type ParserFunc func(contents []byte, url string) ParseResult

//Parser接口
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//请求
type Request struct {
	Url    string //请求的url
	Parser Parser //Parser接口
}

//处理结果
type ParseResult struct {
	Requests []Request //下一步的请求
	Items    []Item    //请求的具体结果
}

type Item struct {
	Url     string //保存url才能去看用户详细信息
	Type    string //Type + Id 避免重复用户 例如：真爱+id:1 世界佳缘+id:1
	Id      string
	Payload interface{}
}

//NilParser
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

//FuncParser
type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

//工厂函数来创建Parser
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
