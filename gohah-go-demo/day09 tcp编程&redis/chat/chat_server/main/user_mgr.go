package main

import "github.com/reminance/reminance-go/gohah-go-demo/day09 tcp编程&redis/chat/chat_server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}
