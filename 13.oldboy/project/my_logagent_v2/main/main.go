package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/my_logagent_v2/config"
	"go_study/13.oldboy/project/my_logagent_v2/etcd"
	"go_study/13.oldboy/project/my_logagent_v2/kafka"
	"go_study/13.oldboy/project/my_logagent_v2/log"
	"go_study/13.oldboy/project/my_logagent_v2/server"
	"go_study/13.oldboy/project/my_logagent_v2/tailf"
)

//日志主程序
func main() {
	//1.conf
	err := config.LoadConf("ini", config.Filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v", err)
		panic("load conf failed")
		return
	}

	//logs.Debug("load conf succ, config:%v", config.AppConfig)

	//2.logger
	err = log.InitLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v", err)
		panic("load logger failed")
		return
	}

	//3.从etcd中读取collect数据
	collectConf, err := etcd.InitEtcd(config.AppConfig.EtcdAddr, config.AppConfig.EtcdKey)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	logs.Debug("init etcd succ")

	//4.tailf
	err = tailf.InitTail(collectConf, config.AppConfig.ChanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	//5.kafka
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

}
