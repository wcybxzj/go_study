package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/day11/logagent/kafka"
	"go_study/13.oldboy/day11/logagent/tailf"
	"time"
	//"fmt"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetOneLine()
		err = sendToKafka(msg)
		if err != nil {
			logs.Error("send to kafka failed, err:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKafka(msg *tailf.TextMsg) (err error) {
	fmt.Printf("read msg:%s, topic:%s\n", msg.Msg, msg.Topic)
	err = kafka.SendToKafka(msg.Msg, msg.Topic)
	return
}
