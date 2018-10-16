package mock

import "fmt"

type Retriever struct {
	Contents string
}

//这个类型实现了fmt.Stringer借口的String方法
//好处是可以在fmt.Println()时候直接打印自定义内容
func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents=%s}", r.Contents)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
