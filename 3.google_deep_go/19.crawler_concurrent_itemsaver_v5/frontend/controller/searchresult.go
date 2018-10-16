package controller

import (
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/frontend/view"
	"gopkg.in/olivere/elastic.v5"

	"context"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/engine"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/frontend/model"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

//locahost:8888/search?q=男 已购房 Payload.Age:(<30)&from=0(分页)
//locahost:8888/search?q=男 已购房 Payload.Age:(<30)&from=10
//locahost:8888/search?q=男 已购房 Payload.Age:(<30)&from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0 //相当于忽略错误
	}
	//fmt.Fprintf(w, "q=%s, from=%d",q, from)

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
/*
搜索:
男 已购房 Age:(<30) Height:(>180)
to
男 已购房 Payload.Age:(<30) Payload.Height:(>180)
*/
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
