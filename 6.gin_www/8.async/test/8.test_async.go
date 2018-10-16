package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"time"
)

const MAX_GOROUTINES = 1000
const EVERY_SIZE = 1000

//const URL = "http://192.168.1.177/order.php?page=%d&size=%d"
const URL = "http://192.168.1.177:8080/long_async/%d"

//const URL ="https://test2-api.nqyong.com/api/order/api/orderListExport?page=%d&size=%d&order_status=&pay_type=&begin_time=&end_time=&order_appid=&visit_id=&kw_type=&keywords="

var db *sql.DB

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//查询当前协程总数
func goroutine_num() {
	go func() {
		for {
			num := runtime.NumGoroutine()
			fmt.Println("goroutine 总数:", num)
			time.Sleep(time.Second)
		}
	}()
}

type worker struct {
	id      int
	page    int
	size    int
	channel chan string
	done    chan bool
}

func workerHttp(workerEntry worker) {
	//url := fmt.Sprintf(URL, workerEntry.page, workerEntry.size)
	//log.Println(url)

	url := fmt.Sprintf(URL, workerEntry.page)
	log.Println(url)

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	workerEntry.channel <- string(body)

	workerEntry.done <- true
}

func createWorker(id int) worker {
	channel := make(chan string)
	done := make(chan bool)
	worker := worker{id, id + 1, EVERY_SIZE, channel, done}
	return worker
}

//请求方式:
//http://192.168.1.177:1234/get_http
//计算总时间
//限制并发总数
func main() {
	goroutine_num()

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/get_http", func(c *gin.Context) {
		var workers [MAX_GOROUTINES]worker
		var results [MAX_GOROUTINES]string

		for i := 0; i < MAX_GOROUTINES; i++ {
			workers[i] = createWorker(i)
		}

		go func() {
			for i := 0; i < MAX_GOROUTINES; i++ {
				str := <-workers[i].channel
				results[i] = str
			}
		}()

		for _, worker := range workers {
			go workerHttp(worker)

		}

		for _, worker := range workers {
			re := <-worker.done
			log.Println("finish", re)
		}

		var final_result string
		for _, result := range results {
			final_result += result
		}

		c.String(http.StatusOK, "%s", final_result)
	})

	r.Run(":5678")
}
