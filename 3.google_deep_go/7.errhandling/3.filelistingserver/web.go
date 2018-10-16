package main

import (
	"log"
	"net/http"
	"os"

	//"github.com/gpmgo/gopm/modules/log"
	"go_study/3.google_deep_go/7.errhandling/3.filelistingserver/filelisting"
)

//统一出错处理
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

//函数式编程:
//函数既可以做参数也可以做返回值
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling"+
				" request:%s", err.Error()) //真实的错误描述
			code := http.StatusOK
			switch {
			//错误类型是文件不存在
			case os.IsNotExist(err):
				code = http.StatusNotFound //404
			case os.IsPermission(err):
				code = http.StatusForbidden //403
			default:
				code = http.StatusInternalServerError //500
			}
			http.Error(
				writer,
				http.StatusText(code), //模糊的错误描述
				code)
		}
	}
}

//显示文件的web server
func main() {
	http.HandleFunc("/list/",
		errWrapper(filelisting.HandleFileListing))
	//addr, errhandler
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
