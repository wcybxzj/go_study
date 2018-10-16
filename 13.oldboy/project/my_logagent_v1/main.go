package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/my_logagent_v1/config"
	"go_study/13.oldboy/project/my_logagent_v1/kafka"
	"go_study/13.oldboy/project/my_logagent_v1/log"
	"go_study/13.oldboy/project/my_logagent_v1/server"
	"go_study/13.oldboy/project/my_logagent_v1/tailf"
	"time"
)

func main() {
	//conf
	filename := "/root/www/go_www/src/go_study/13.oldboy/project/my_logagent_v1/logagent.conf"
	err := config.LoadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v", err)
		panic("load conf failed")
		return
	}

	//logs.Debug("load conf succ, config:%v", config.AppConfig)

	//logger
	err = log.InitLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v", err)
		panic("load logger failed")
		return
	}

	//tailf
	err = tailf.InitTail(config.AppConfig.CollectConf, config.AppConfig.ChanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	//kafka
	err = kafka.InitKafka(config.AppConfig.KafkaAddr)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}

	logs.Debug("init all succ")
	err = server.ServerRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}
	logs.Info("program exit")

	time.Sleep(time.Second)
}
