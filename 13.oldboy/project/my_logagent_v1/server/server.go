package server

import (
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/my_logagent_v1/kafka"
	"go_study/13.oldboy/project/my_logagent_v1/tailf"
	"time"
)

func ServerRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err := kafka.SendToKafka(msg.Msg, msg.Topic)
		if err != nil {
			logs.Error("send kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}
