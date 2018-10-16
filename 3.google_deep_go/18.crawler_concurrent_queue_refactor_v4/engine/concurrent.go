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
//v4 重构第1步
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

//把这个函数单独抽成一个借的用处
//createWorker()传递的时候只传递 ReadyNotifier即可，不用传Scheduler
type ReadyNotifier interface{
	WorkerReady ( chan Request)
}

func (e *ConcurrentEngine)Run (seeds ...Request)  {
	out := make(chan ParserResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount ;i++  {
		//v4 重构第3步
		createWorker(e.Scheduler.WorkerChan(), out , e.Scheduler)
	}

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount := 0 //统计获取了多少item
	//engine的任务是分发，不要做耗时的工作，获取items或者requests要立刻交出去
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

//v4 重构第2步
func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier)  {
	go func() {
		for  {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

