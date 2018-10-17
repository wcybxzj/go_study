package router

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/seckill/day14/SecProxy/controller"
)

func init() {
	logs.Debug("enter router init")
	//url,object,method
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}
