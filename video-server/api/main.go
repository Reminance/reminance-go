package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	fmt.Printf("server listening at 8000... time: %d", time.Now().Unix())
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Errorf("err lietening at 8000, err: %s", err.Error())
	}
}