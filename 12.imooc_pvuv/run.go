package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

//需求模拟产生 nginx access.log

//首页:http://dz3.com/forum.php
//列表:http://dz3.com/forum-2-1.html	 	--->  forum-{fid}-{page}.html
//详情页:http://dz3.com/thread-1-1-1.html --->  thread-{tid}-{page}-{prevpage}.html
var ualist = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36 OPR/26.0.1656.60",
	"Opera/8.0 (Windows NT 5.1; U; en)",
	"Mozilla/5.0 (Windows NT 5.1; U; en; rv:1.8.1) Gecko/20061208 Firefox/2.0.0 Opera 9.50",
	"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; en) Opera 9.50",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:34.0) Gecko/20100101 Firefox/34.0",
	"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/30.0.1599.101 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
}

type resource struct {
	url    string
	target string
	start  int
	end    int
}

func ruleResource() []resource {
	var res []resource
	r1 := resource{
		url:    "http://dz3.com/forum.php",
		target: "",
		start:  0,
		end:    0,
	}

	r2 := resource{
		url:    "http://dz3.com/forum-{$id}-1.html",
		target: "{$id}",
		start:  1,
		end:    21,
	}

	r3 := resource{
		url:    "http://dz3.com/thread-{$id}-1-1.html",
		target: "{$id}",
		start:  1,
		end:    12000,
	}

	res = append(res, r1)
	res = append(res, r2)
	res = append(res, r3)

	return res
}

func buildUrl(res []resource) []string {
	var list []string

	for _, resItem := range res {
		if len(resItem.target) == 0 {
			list = append(list, resItem.url)
		} else {
			for i := resItem.start; i <= resItem.end; i++ {
				urlStr := strings.Replace(resItem.url, resItem.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

/*
log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
			  '$status $body_bytes_sent "$http_referer" '
			  '"$http_user_agent" "$http_x_forwarded_for"';
*/
func makeLog(current, refer, ua string) string {
	//模拟tongji.js上传的GET中的数据
	u := url.Values{}
	u.Set("time", "1")    //js采集到的用户访问的客户端时间
	u.Set("url", current) //js采集到的用户访问的url
	u.Set("refer", refer) //js采集到的用户访问的refer url
	u.Set("ua", ua)       //js采集到的用户访问的浏览器类型
	paramsStr := u.Encode()

	log := "127.0.0.1 - - [16/Sep/2018:21:14:30 +0800]" +
		" \"GET /dig?{$paramsStr} HTTP/1.1\" 200 43 \"http://dz3.com/thread-1-2-1.html\" " +
		"\"{$ua}\" \"-\""

	log = strings.Replace(log, "{$paramsStr}", paramsStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)
	return log
}

func randInt(min, max int) int {
	//rand.New(随机种子)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max {
		return max
	}
	//Intn(num)==>return a radom number[0,num]
	return r.Intn(max-min) + min
}

// go run run.go --total=300000 --filePath=/usr/local/nginx/logs/dig.log
func main() {
	total := flag.Int("total", 100, "how many rows by created")
	filePath := flag.String("filePath", "/usr/local/nginx/logs/dig.log", "log file path")
	flag.Parse()

	fmt.Println(*total, *filePath)

	//构造访问的url
	res := ruleResource()
	list := buildUrl(res)

	//for _,r := range list{
	//	fmt.Println(r)
	//}

	////根据url, 生成制定行的日志
	var logStr string
	for i := 0; i <= (*total); i++ {
		currentUrl := list[randInt(0, len(list)-1)]
		referUrl := list[randInt(0, len(list)-1)]
		ua := ualist[randInt(0, len(ualist)-1)]

		logStr = logStr + makeLog(currentUrl, referUrl, ua) + "\n"

		//覆盖写
		//ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	}

	fd, err := os.OpenFile(*filePath, (os.O_CREATE | os.O_RDWR | os.O_APPEND), 0644)
	if err != nil {
		panic(err)
	}
	fd.Write([]byte(logStr))
	fd.Close()
	fmt.Println("done.\n")
}
