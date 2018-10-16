package scheduler

import "go_study/3.google_deep_go/17.crawler_concurrent_queue_v3/engine"

//简单调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

//冲突版:
//func (s *SimpleScheduler) Submit(r engine.Request) {
//	//send request down to worker chan
//	s.workerChan <- r
//}

//解决版:
func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func(){s.workerChan <- r}()
}