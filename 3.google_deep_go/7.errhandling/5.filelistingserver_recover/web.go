package main

import (
	"log"
	"net/http"
	"os"

	"go_study/3.google_deep_go/7.errhandling/5.filelistingserver_recover/filelisting"
	//"github.com/gpmgo/gopm/modules/log"
)

//统一出错处理
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

//函数式编程:
//函数既可以做参数也可以做返回值
func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		//错误处理1:
		//处理一下pannic,使用defer+recover
		//访问:http://192.168.91.15:8888/abc
		//终端默认:一大堆stack trace信息
		//我们recover这样就不用让/usr/local/go/src/net/http/server.go:1726 进行recvoer
		defer func() {
			r := recover()
			if r != nil { //如果recover返回nil没必要做处理
				log.Printf("Panic :%v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("handler返回err:%s", err.Error()) //真实的错误描述
			//错误处理2:
			//用户自定义错误,用Type Assertion识别出用户定义error
			//访问: http://192.168.91.15:8888/list1
			//如果是userErr借口类型的err可以直接返回给用户
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest) //http 400
				return
			}

			//错误处理3:
			//正常的出错处理,一般是通用的错误
			//访问:http://192.168.91.15:8888/list/123
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

type userError interface {
	error            //借口组合, builtin/error.go中的error借口
	Message() string //给用户看的error
}

//显示文件的web server
func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileListing))
	//addr, errhandler
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
