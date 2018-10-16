package fetcher

import (
	"net/http"
	"io/ioutil"
		"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"fmt"

	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/unicode"
	"log"
)

//功能:就是获取html,并且转换成utf-8编码
func Fetch(url string)  ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//println("Error: status code", resp.StatusCode)

		//create user custom err func 1:
		//return nil, errors.New("http Status is not 200")

		//create user custom err func 2:
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	//问题:出现乱码
	//方法1:手动指定原来的编码
	//通过自己观察网页<meta charset="gbk" />
	//将传入的reader(resp.Body)转换成新的reader
	//utf8Reader := transform.NewReader(resp.Body,
	//	simplifiedchinese.GBK.NewDecoder())

	//方法2:自动识别编码类型
	//from  encoding get decoder
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//1.html转码 gbk->utf8
//go get -v golang.org/x/text
//或者
//go get -v golang.org/x/net

//2.根据网站编码决定如何转换
//方法1:
// <meta charset="gbk" /> 获取里面的gbk,
// 缺点是也许它说是gbk但是页面却不是gbk
//方法2(推荐):
// 用一个golang的库来识别页面编码,来保证准确
// go get -v golang.org/x/net
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//Peek只是引用前1024bytes 放入bufio,并没有去移动reader不会影响后边的读取
	bytes, err := r.Peek(1024)

	//if peek fail, return default encoding
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	//put bytes into DetermineEncoding, get charset encoding
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
