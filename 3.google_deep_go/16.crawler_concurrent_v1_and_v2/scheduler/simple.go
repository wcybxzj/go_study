package scheduler

import "go_study/3.google_deep_go/16.crawler_concurrent_v1_and_v2/engine"

//简单调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

//并发爬虫第一版本:简单调度器(失败)
//func (s *SimpleScheduler) Submit(r engine.Request) {
//	//send request down to worker chan
//	s.workerChan <- r
//}

//并发爬虫第二版本:并发调度器(成功)
func (s *SimpleScheduler) Submit(r engine.Request) {
	//send request down to worker chan
	go func(){s.workerChan <- r}()
}