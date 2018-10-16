package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mgutz/str"
	"github.com/sirupsen/logrus"
	"io"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
http://dz3.com/
http://dz3.com/forum-2-1.html
http://dz3.com/thread-11111-1-1.html
*/
const HANDLE_DIG = " /dig?"
const HANDLE_THREAD = "/thread-"
const HANDLE_FORUM = "/forum-"
const HANDLE_THREAD_HTML = "-1-1.html"
const HANDLE_FORUM_HTML = "-1.html"

type cmdParams struct {
	logFilePath string
	routineNum  int
}

//tongji.js:GET发来的数据
type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data  digData
	uid   string
	unode urlNode
}

type urlNode struct {
	unType string //urlnode的type: 首页(index) 列表(forum)页或者内容页(thread)
	unRid  int    //Resource id: 资源id
	unUrl  string //当前这个页面的url
	unTime string //当前访问的时间
}

//存储结构体
type storageBlock struct {
	counterType  string
	storageModel string
	unode        urlNode
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

//go run my_analysis.go -l /tmp/test.log

//redis连接池:
//1.在高并发推荐使用连接池
//2.在无人访问时候链接池会断开, 所以要定时的去ping一下防止断开
func main() {
	//获取参数
	logFilePath := flag.String(
		"logFilePath",
		"/usr/local/nginx/logs/dig.log",
		"log file path")

	routineNum := flag.Int(
		"routineNum", 5,
		"consumer goroutine num")

	l := flag.String(
		"l",
		"/tmp/log",
		"program log")
	flag.Parse()

	params := cmdParams{*logFilePath, *routineNum}

	//本程序日志
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	log.Out = logFd
	defer logFd.Close()

	log.Infof("start!")
	log.Infof("Params: logFilePath=%s, routineNum:%d", params.logFilePath, params.routineNum)

	//初始化channel, 用于数据传递
	var logChannel = make(chan string, 3*params.routineNum) //这个channel的数据量比较大
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)

	redisPool, err := pool.New("tcp", "localhost:6379", params.routineNum*2)
	if err != nil {
		logrus.Fatalln("Redis pool created failed.")
		panic(err)
	} else {
		go func() {
			redisPool.Cmd("PING")
			time.Sleep(3 * time.Second)
		}()
	}

	//日志消费者 goroutine
	go readFileLinebyLine(params, logChannel)

	//创建一组日志解析 goroutines
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	//pv uv 统计器 goroutines
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel, redisPool)
	//可扩展xxxCounter

	//存储器 goroutine
	go dataStorage(storageChannel, redisPool)

	time.Sleep(time.Second * 1000)
}

/*
1.加洋葱皮:
加洋葱皮有2类:页面结构和时间
存储方式:redis SortedSet

1.1 时间:
天-小时-分钟
是说一个用户的访问当前的分钟,小时,天 都要加1

1.2 页面结构:
例如一个论坛
首页->列表页->详情页
用户访问详情页,列表页和首页也要加1
例如一个商城
首页->大分类->小分类->内容页
用户访问详情页,小分类,大分类,首页都要加1

2.数据量:
原始数据 * 统计协程数 * 洋葱皮的数量
原始数据:1W
统计协程数:pv和uv 2
洋葱皮的数量:6

1W * 2 *6 =12W


3.存储方式:
大企业:推荐HBase,需要列提前定义好,列比较固定
小企业:推荐redis
*/
func dataStorage(storageChannel chan storageBlock, redisPool *pool.Pool) {
	for block := range storageChannel {
		prefix := block.counterType + "_"

		//洋葱皮
		setKeys := []string{
			prefix + "day_" + getTime(block.unode.unTime, "day"),
			prefix + "hour_" + getTime(block.unode.unTime, "hour"),
			prefix + "min_" + getTime(block.unode.unTime, "min"),
			prefix + block.unode.unType + "_day_" + getTime(block.unode.unTime, "day"),
			prefix + block.unode.unType + "_hour_" + getTime(block.unode.unTime, "hour"),
			prefix + block.unode.unType + "_min_" + getTime(block.unode.unTime, "min"),
		}

		rowId := block.unode.unRid

		for _, key := range setKeys {
			ret, err := redisPool.Cmd(block.storageModel, key, 1, rowId).Int()
			if ret <= 0 || err != nil {
				//因为有少量的错误不该退出,在统计时候可以忽略一些错误
				log.Errorln("DataStorage redis storage error.",
					block.storageModel, key, rowId)
			}
		}
	}
}

//pv:网站访问的请求量
//用redis的sortedset来存储数据
func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {
	for data := range pvChannel {
		sItem := storageBlock{"pv", "ZINCRBY", data.unode}
		storageChannel <- sItem
	}
}

//uv:网站的访问用户数
//需要对uid进行去重使用redis 中 HyperLoglog位图来进行去重
func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock, redisPool *pool.Pool) {
	for data := range uvChannel {
		//uv一般是按天去重
		hyperLoglogKey := "uv_hpll_" + getTime(data.data.time, "day")
		ret, err := redisPool.Cmd("PFADD", hyperLoglogKey, data.uid, "EX", 86400).Int()
		if err != nil {
			log.Warnln("uvCounter check redis hyperloglog failed, err:", err)
		}
		//添加失败,说明这个value已经存在
		if ret != 1 {
			continue
		}

		sItem := storageBlock{"uv", "ZINCRBY", data.unode}
		storageChannel <- sItem
	}
}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) error {
	for logStr := range logChannel {
		//切割日志字符串,扣出打点上报的数据
		data := cutLogFetchData(logStr)

		//uid:正常的网站哪怕不登录也有一个id在cookie中
		//这个列子访问数据是自己创建的所以用md5(refer+ua)来模拟uid
		//用uid进行去重
		hasher := md5.New()
		hasher.Write([]byte(data.refer + data.ua))
		uid := hex.EncodeToString(hasher.Sum(nil))

		//很多解析工作可以在这里写
		//....
		//....

		uData := urlData{data, uid, formatUrl(data.url, data.time)}

		//log.Infoln(uData)

		pvChannel <- uData
		uvChannel <- uData
	}
	return nil
}

/*
"GET /dig?refer=aaaa&time=1&ua=bbbb&url=cccc HTTP/1.1"
*/
func cutLogFetchData(logStr string) digData {
	strings.TrimSpace(logStr)
	pos1 := str.IndexOf(logStr, HANDLE_DIG, 0)
	if pos1 == -1 {
		return digData{}
	}
	pos1 += len(HANDLE_DIG)
	pos2 := str.IndexOf(logStr, "HTTP/", pos1)
	if pos2 == -1 {
		return digData{}
	}

	d := str.Substr(logStr, pos1, pos2-pos1)
	//解析进行编码的url,必须加上http://xxxx/?,否则无法工作
	urlInfo, err := url.Parse("http://localhost/?" + d)
	if err != nil {
		return digData{}
	}
	data := urlInfo.Query()
	return digData{
		data.Get("time"),
		data.Get("refer"),
		data.Get("url"),
		data.Get("ua"),
	}
}

func readFileLinebyLine(params cmdParams, logChannel chan string) error {
	fd, err := os.Open(params.logFilePath)
	if err != nil {
		log.Warningf("readFileLinebyLine can not open file:%s", params.logFilePath)
		return err
	}
	defer fd.Close()

	count := 0
	bufferReader := bufio.NewReader(fd)
	for {
		line, err := bufferReader.ReadString('\n')
		//fmt.Println(line)
		logChannel <- line //写入channel
		count++

		if count%(1000*params.routineNum) == 0 {
			log.Infoln("readFileLinebyLine line:%d", count)
		}

		if err != nil {
			if err == io.EOF {
				time.Sleep(3 * time.Second)
				log.Infof("readFileLinebyLine wait, count:%d", count)
			} else {
				log.Warningf("ReadFileLinebyLine read log error")
			}
		}
	}
	return nil
}

func formatUrl(url, time string) urlNode {
	//一定从访问量大的着手,详情页>列表页>=首页
	pos1 := str.IndexOf(url, HANDLE_THREAD, 0)
	if pos1 != -1 {
		pos1 += len(HANDLE_THREAD)
		pos2 := str.IndexOf(url, HANDLE_THREAD_HTML, 0)
		idStr := str.Substr(url, pos1, pos2-pos1)
		id, _ := strconv.Atoi(idStr)
		return urlNode{"thread", id, url, time}
	} else {
		pos1 = str.IndexOf(url, HANDLE_FORUM, 0)
		if pos1 != -1 {
			pos1 += len(HANDLE_FORUM)
			pos2 := str.IndexOf(url, HANDLE_FORUM_HTML, 0)
			idStr := str.Substr(url, pos1, pos2-pos1)
			id, _ := strconv.Atoi(idStr)
			return urlNode{"forum", id, url, time}
		} else {
			return urlNode{"index", 1, url, time}
		}
	}
	return urlNode{}
}

//根据类型产生10进制timestamp字符串
func getTime(logTime, timeType string) string {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"

	case "hour":
		item = "2006-01-02 15"

	case "min":
		item = "2006-01-02 15:04"
	}
	//根据结构产生time struct
	time, _ := time.Parse(item, time.Now().Format(item))
	//10进制timestamp 转成 10进制字符串
	return strconv.FormatInt(time.Unix(), 10)
}
