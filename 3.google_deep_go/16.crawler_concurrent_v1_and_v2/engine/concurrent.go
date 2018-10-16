package engine

import (
	"log"
)

//并发版engine
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

//调度器
//声明的位置有掉别扭
type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine)Run (seeds ...Request)  {
	in := make(chan Request)
	out := make(chan ParserResult)

	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount ;i++  {
		createWorker(in, out)
	}

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount := 0 //统计获取了多少item
	for {
		result := <-out //冲突点-----------------------------
		for _, item := range result.Items {
			log.Printf("Got item count:%d, item: %v",
				itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)//冲突点--------------
		}
	}
}

func createWorker(in chan Request, out chan ParserResult)  {
	go func() {
		for  {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

