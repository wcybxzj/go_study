package engine

type ParserFunc func(contents []byte, url string) ParserResult

//请求
type Request struct {
	Url        string     //请求的url
	ParserFunc ParserFunc //处理请求的函数
}

//处理结果
type ParserResult struct {
	Requests []Request //下一步的请求
	Items    []Item    //请求的具体结果
}

type Item struct {
	Url     string //保存url才能去看用户详细信息
	Type    string //Type + Id 避免重复用户 例如：真爱+id:1 世界佳缘+id:1
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
