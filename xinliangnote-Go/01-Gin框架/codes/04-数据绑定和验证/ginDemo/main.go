package main

import (
	"fmt"
	"github.com/reminance/reminance-go/xinliangnote-Go/01-Gin框架/codes/04-数据绑定和验证/ginDemo/config"
	"github.com/reminance/reminance-go/xinliangnote-Go/01-Gin框架/codes/04-数据绑定和验证/ginDemo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.New()
	router.InitRouter(engine) // 设置路由
	err := engine.Run(config.PORT)
	if err != nil {
		fmt.Println(err.Error())
	}
}
