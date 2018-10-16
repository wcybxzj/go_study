package main

import (
	_ "go_study/13.oldboy/day13/beego_example/router"
	"github.com/astaxie/beego"
	"fmt"
)


// cd /root/www/go_www/src/go_study/13.oldboy/day13/beego_example
// go build -o beggo main/*
// ./beggo
// 访问 127.0.0.1:9091/index
func main() {
	fmt.Println("main")
	beego.Run()
}
