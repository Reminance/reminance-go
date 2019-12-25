package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//定义middleware 用于在router之前做通用处理  如鉴权
type middleWareHandler struct {
	router *httprouter.Router
}

func newMiddleWareHandler(router *httprouter.Router) http.Handler {
	handler := middleWareHandler{}
	handler.router = router
	return handler
}

func (handler middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//通用处理
	//check session
	validateUserSession(r)
	handler.router.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	middleWareHandler := newMiddleWareHandler(r)
	fmt.Printf("server listening at 8000... time: %d", time.Now().Unix())
	err := http.ListenAndServe(":8000", middleWareHandler)
	if err != nil {
		fmt.Errorf("err lietening at 8000, err: %s", err.Error())
	}
}
