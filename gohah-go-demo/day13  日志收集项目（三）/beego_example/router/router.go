package router

import (
	"github.com/reminance/reminance-go/gohah-go-demo/day13  日志收集项目（三）/beego_example/controller/IndexController"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
