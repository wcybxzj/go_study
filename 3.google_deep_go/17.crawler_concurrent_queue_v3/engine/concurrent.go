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
	WorkerReady ( chan Request)
	Run()
}

func (e *ConcurrentEngine)Run (seeds ...Request)  {
	out := make(chan ParserResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount ;i++  {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount := 0 //统计获取了多少item
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item count:%d, item: %v",
				itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(out chan ParserResult, s Scheduler)  {
	in := make(chan Request)
	go func() {
		for  {
			//worker tell scheduler i'm ready
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

