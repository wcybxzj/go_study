package IndexController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController2 struct {
	beego.Controller
}

func (p *IndexController2) Index2() {
	logs.Debug("enter index2 controller")
	p.Data["json"] = 200
	p.ServeJSON(true)
}
