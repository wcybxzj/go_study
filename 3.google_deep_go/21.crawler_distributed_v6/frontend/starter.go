package main

import (
	"go_study/3.google_deep_go/21.crawler_distributed_v6/frontend/controller"
	"net/http"
)

const view_path = `3.google_deep_go/21.crawler_distributed_v6/frontend/view`

// http://localhost:8888/search
// http://localhost:8888/search?q=男 已购车 已购房&from=67

//正式搜索：
//http://localhost:8888/
//搜索框:男 已购房 Age:(<30) Height:(>180)
//搜索框:女 Age:(<25) Height:(>165) Weight:([1 TO 50])
func main() {
	http.Handle("/",
		http.FileServer(http.Dir(view_path)))
	http.Handle("/search",
		controller.CreateSearchResultHandler(view_path+"/template.html"))

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		panic(err)
	}
}
