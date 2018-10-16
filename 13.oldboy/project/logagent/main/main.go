package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/logagent/kafka"
	"go_study/13.oldboy/project/logagent/tailf"
	//"time"
)

func main() {
	//1.conf
	filename := "/root/www/go_www/src/go_study/13.oldboy/day11/logagent/logagent.conf"
	err := loadConf("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	//2.logger
	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("load conf succ, config:%v", appConfig)

	collectConf, err := initEtcd(appConfig.etcdAddr, appConfig.etcdKey)
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	logs.Debug("initialize etcd succ")

	err = tailf.InitTail(collectConf, appConfig.chanSize)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	logs.Debug("initialize tailf succ")
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init tail failed, err:%v", err)
		return
	}

	logs.Debug("initialize all succ")
	/*
		go func() {
			var count int
			for {
				count++
				logs.Debug("test for logger %d", count)
				time.Sleep(time.Millisecond * 1000)
			}
		}()*/
	err = serverRun()
	if err != nil {
		logs.Error("serverRUn failed, err:%v", err)
		return
	}

	logs.Info("program exited")
}
