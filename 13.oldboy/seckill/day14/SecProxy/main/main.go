package main

import (
	"github.com/astaxie/beego"
	"go_study/13.oldboy/seckill/day14/SecProxy/conf"
	_ "go_study/13.oldboy/seckill/day14/SecProxy/router"
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
