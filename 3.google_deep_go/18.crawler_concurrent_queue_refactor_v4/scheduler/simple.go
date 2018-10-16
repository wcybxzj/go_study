package scheduler

import "go_study/3.google_deep_go/18.crawler_concurrent_queue_refactor_v4/engine"

//简单调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

//v4 重构第5步
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

//v4 重构第6步
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

//v4 重构第7步
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
//	s.workerChan = c
//}

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