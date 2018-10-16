package server

import (
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/my_logagent_v2/kafka"
	"go_study/13.oldboy/project/my_logagent_v2/tailf"
	"time"
)

func ServerRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		//fmt.Println("msg:", msg.Msg)
		err := kafka.SendToKafka(msg.Msg, msg.Topic)
		if err != nil {
			logs.Error("send kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
