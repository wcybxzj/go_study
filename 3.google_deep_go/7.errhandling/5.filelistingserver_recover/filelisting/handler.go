package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

//
type userError2 string

//实现builtin/builtin.go 中的error借口
func (e userError2) Error() string {
	return "::::userError2.Error()::::" + string(e)
}

//实现web.go中的userError借口
func (e userError2) Message() string {
	return "::::userError2.Message()::::" + string(e)
}

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 { //只要不是/list/开头就报错
		return userError2("Path must start with " + prefix)
	}
	path := request.URL.Path[len(prefix):] // /list/fib.txt
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
