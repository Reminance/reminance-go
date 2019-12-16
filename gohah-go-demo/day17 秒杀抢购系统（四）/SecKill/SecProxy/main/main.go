package main

import (
	_ "github.com/reminance/reminance-go/gohah-go-demo/day17 秒杀抢购系统（四）/SecKill/SecProxy/router"

	"github.com/astaxie/beego"
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
