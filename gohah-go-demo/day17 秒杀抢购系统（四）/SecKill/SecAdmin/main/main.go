package main

import (
	"github.com/astaxie/beego"
	_ "github.com/reminance/reminance-go/gohah-go-demo/day17 秒杀抢购系统（四）/SecKill/SecAdmin/router"
	"fmt"
)

func main() {
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
