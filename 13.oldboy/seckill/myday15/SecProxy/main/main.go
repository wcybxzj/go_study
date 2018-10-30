package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_study/13.oldboy/seckill/myday15/SecProxy/conf"
	_ "go_study/13.oldboy/seckill/myday15/SecProxy/router"
)

func main() {
	fmt.Println("work!")

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
