package router

import (
	"github.com/reminance/reminance-go/gohah-go-demo/day16 秒杀抢购系统（三）/SecKill/SecProxy/controller"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.Debug("enter router init")
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}
