package view

import (
	"html/template"
	"io"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return  SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}
}

func (s SearchResultView) Render (w io.Writer, data model.SearchResult) error{
	return s.template.Execute(w, data)
}
