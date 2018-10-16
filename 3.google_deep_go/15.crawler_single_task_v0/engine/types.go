package engine

//请求
type Request struct{
	Url string //请求的url
	ParserFunc func([]byte)ParserResult //处理请求的函数
}

//处理结果
type ParserResult struct{
	Requests []Request //下一步的请求
	Items []interface{} //请求的具体结果
}

func NilParser([]byte)  ParserResult {
	return ParserResult{}
}