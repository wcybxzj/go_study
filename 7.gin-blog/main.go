package main

import (
	"fmt"
	"net/http"

	"go_study/7.gin-blog/pkg/setting"
	"go_study/7.gin-blog/routers"
)

//http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test
func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
