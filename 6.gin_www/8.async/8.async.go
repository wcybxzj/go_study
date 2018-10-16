package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

const MAX_GOROUTINES = 100

var WorkerChan chan int

var JobChan chan int

type Job struct {
	JobId string `json:"job_id" binding:"required"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Exec  string `json:"exec"`
}

const STATE_IDLE = 0
const STATE_BUSY = 1

type Worker struct {
	workId int
	in     chan int
	state  int
}

////空闲的worker加入到channel
//func createWorkerConnectionPool(num int) chan Worker {
//	var clients []*rpc.Client
//	for _, h := range hosts {
//		client, err := rpcsupport.NewClient(h)
//		if err == nil {
//			clients = append(clients, client)
//			log.Printf("Connected to %s", h)
//		} else {
//			log.Printf("Error connecting to %s: %v", h, err)
//		}
//	}
//
//	out := make(chan *rpc.Client)
//
//	go func() {
//		for {
//			for _, client := range clients {
//				out <- client
//			}
//		}
//	}()
//	return out
//}

func doWork(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second * 3)
		fmt.Printf("worker %d received %c\n", id, n)
	}
}

func createWorker(id int) Worker {
	w := Worker{
		workId: id,
		in:     make(chan int),
		state:  STATE_IDLE,
	}
	go doWork(id, w.in)
	return w
}

//http://192.168.1.177:8080/long_async/123
func main() {
	r := gin.Default()
	//1.协程池
	var workers [MAX_GOROUTINES]Worker
	for i := 0; i < MAX_GOROUTINES; i++ {
		workers[i] = createWorker(i)
	}

	//worker_schedule_pool
	//workerSchedulePool(MAX_GOROUTINES)

	//2.异步
	r.GET("/long_async/:id", func(c *gin.Context) {
		id_int := c.Param("id")

		id, err := strconv.Atoi(id_int)
		if err != nil {
			panic(err)
		}

		// goroutine 中只能使用只读的上下文 c.Copy()
		cCp := c.Copy()
		go func() {

			fmt.Println(id)
			//time.Sleep(5 * time.Second)
			// 注意使用只读上下文
			log.Printf("Done! in path：%s id:%d"+cCp.Request.URL.Path, id)

		}()
		c.JSON(http.StatusOK, gin.H{"status": "finish"})
	})

	r.Run(":8080")
}
