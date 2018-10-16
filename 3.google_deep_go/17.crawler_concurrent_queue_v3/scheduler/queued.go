package scheduler

import "go_study/3.google_deep_go/17.crawler_concurrent_queue_v3/engine"

type QueuedScheduler struct{
	requestChan chan engine.Request
	workerChan chan chan engine.Request //workerChan字段的类型是chan, 值是chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady (w chan engine.Request){
	s.workerChan <- w
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

//功能:
//1.创建workerChan和requestChan
//2.启动一个协程,等待worker或者request到来
func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for  {
			//只要用了select所有channel操作都放select里
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			//scheduler收到request或者worker需要排队
			select {
			case r := <-s.requestChan:
				requestQ =append(requestQ, r)

			case w := <-s.workerChan:
				workerQ =append(workerQ, w)

			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}