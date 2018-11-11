package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"

	VIDEO_PATH = "./videos/"
)

//生产者/消费者 通信用的channel
type controlChan chan string

//生产者/消费者 传递task的channel
type dataChan chan interface{}

//生产者/消费者
type fn func(dc dataChan) error
