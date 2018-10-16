package tailf

import (
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"go_study/13.oldboy/project/my_logagent_v2/config"
	"sync"
	"time"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type TailObj struct {
	tail     *tail.Tail
	conf     config.CollectConf
	status   int
	exitChan chan int
}

type TextMsg struct {
	Msg   string
	Topic string
}

//tailf对象管理类
type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
}

var (
	tailObjMgr *TailObjMgr
)

//从channel获取数据
func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return
}

func InitTail(conf []config.CollectConf, chanSize int) (err error) {
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	if len(conf) == 0 {
		//err = fmt.Errorf("invalid config for log collect,conf:%v", conf)
		logs.Error("invalid config for log collect,conf:%v", conf)
		return
	}

	for _, v := range conf {
		createNewTask(v)
	}

	return nil
}

//如果任务正在执行不做处理
//只有任务不存在,才会把新的日志监控启动一个go去进行监控
func UpdateConfig(confs []config.CollectConf) (err error) {
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()

	for _, oneConf := range confs {
		var isRunning = false
		for _, obj := range tailObjMgr.tailObjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}

		if isRunning {
			continue
		} else {
			createNewTask(oneConf)
		}
	}

	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range confs {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}

		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		tailObjs = append(tailObjs, obj)
	}
	return
}

func createNewTask(conf config.CollectConf) {
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}

	tails, errTail := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if errTail != nil {
		//logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, errTail)
		return
	}

	obj.tail = tails
	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

	go readFromTail(obj)

}

//相当于每个collect启动一个go进行读取
func readFromTail(tailObj *TailObj) {
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines:
			if !ok {
				logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			//fmt.Println("msg:", line)
			textMsg := &TextMsg{
				Msg:   line.Text,
				Topic: tailObj.conf.Topic,
			}

			tailObjMgr.msgChan <- textMsg

		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited, conf:%v", tailObj.conf)
			return
		}
	}
}
