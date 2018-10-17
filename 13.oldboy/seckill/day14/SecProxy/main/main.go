package main

import (
	"github.com/astaxie/beego"
	_ "go_study/13.oldboy/seckill/day14/SecProxy/router"
)

func main() {

	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
		return
	}

	beego.Run()
}
