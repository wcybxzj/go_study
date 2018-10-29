package main

import (
	"github.com/astaxie/beego"
	_ "go_study/13.oldboy/seckill/myday15/SecProxy/router"
	"go_study/13.oldboy/seckill/myday15/SecProxy/conf"
)

func main() {
	err := conf.InitConfig()
	if err != nil {
		panic(err)
		return
	}

	err = conf.InitSec()
	if err != nil {
		panic(err)
		return
	}

	beego.Run()
}
