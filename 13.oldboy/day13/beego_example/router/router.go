package router

import (
	"go_study/13.oldboy/day13/beego_example/controller/IndexController"
	"github.com/astaxie/beego"
)

func init() {
	//*:Index: *的意思是get/post都接收
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
	beego.Router("/index2", &IndexController.IndexController2{}, "*:Index2")
}
