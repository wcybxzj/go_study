package main

import (
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/log_transfer/config"
	"go_study/13.oldboy/project/log_transfer/server"
	"go_study/13.oldboy/project/log_transfer/tools"
)

func main() {
	err := config.InitConfig("ini", config.Filename)
	if err != nil {
		panic(err)
		return
	}
	//fmt.Println(config.LogConfig)

	err = tools.InitLogger(config.LogConfig.LogPath, config.LogConfig.LogLevel)
	if err != nil {
		panic(err)
		return
	}
	logs.Debug("init logger succ")

	err = tools.InitKafka(config.LogConfig.KafkaAddr, config.LogConfig.KafkaTopic)
	if err != nil {
		logs.Error("init kafka failed, err:%v", err)
		return
	}
	logs.Debug("init kafka succ")

	err = tools.InitES(config.LogConfig.ESAddr)
	if err != nil {
		logs.Error("init es failed, err:%v", err)
		return
	}

	logs.Debug("init es client succ")

	err = server.Run()
	if err != nil {
		logs.Error("run  failed, err:%v", err)
		return
	}

	logs.Warn("warning, log_transfer is exited")
}
