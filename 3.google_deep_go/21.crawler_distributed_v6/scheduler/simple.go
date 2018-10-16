package scheduler

import "go_study/3.google_deep_go/21.crawler_distributed_v6/engine"

//简单调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//解决版:
func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func() { s.workerChan <- r }()
}
