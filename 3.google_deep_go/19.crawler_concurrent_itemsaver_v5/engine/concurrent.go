package engine

//并发版engine
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

//调度器
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
		createWorker(e.Scheduler.WorkerChan(), out , e.Scheduler)
	}

	for _, r := range seeds{
		if isDuplicate(r.Url) {
			continue
		}

		e.Scheduler.Submit(r)
	}

	//engine的任务是分发，不要做耗时的工作，获取items或者requests要立刻交出去
	for {
		result := <-out
		for _, item := range result.Items {
			//每一个item都用一个goroutine送给itemsaver
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

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

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}


