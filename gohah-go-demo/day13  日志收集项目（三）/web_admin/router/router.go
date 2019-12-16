package router

import (
	"github.com/reminance/reminance-go/gohah-go-demo/day13  日志收集项目（三）/web_admin/controller/AppController"
	"github.com/reminance/reminance-go/gohah-go-demo/day13  日志收集项目（三）/web_admin/controller/LogController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/list", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/apply", &AppController.AppController{}, "*:AppApply")

	beego.Router("/app/create", &AppController.AppController{}, "*:AppCreate")

	beego.Router("/log/apply", &LogController.LogController{}, "*:LogApply")
	beego.Router("/log/list", &LogController.LogController{}, "*:LogList")
	beego.Router("/log/create", &LogController.LogController{}, "*:LogCreate")

}
