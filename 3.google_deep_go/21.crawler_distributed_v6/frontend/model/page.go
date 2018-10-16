package model

type SearchResult struct {
	Hits     int64 //数据总数
	Start    int   //开始位置
	Query    string
	PrevFrom int
	NextFrom int
	Items    []interface{}
}
