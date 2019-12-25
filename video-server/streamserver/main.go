package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//定义middleware 用于在router之前做通用处理  如流控
type middleWareHandler struct {
	router      *httprouter.Router
	connLimiter *ConnLimiter
}

func newMiddleWareHandler(router *httprouter.Router, concurrentConn int) http.Handler {
	handler := middleWareHandler{}
	handler.router = router
	handler.connLimiter = NewConnLimiter(concurrentConn)
	return handler
}

func (handler middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//通用处理
	//check rate limit
	if !handler.connLimiter.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
		return
	}
	handler.router.ServeHTTP(w, r)
	defer handler.connLimiter.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:video-id", streamHandler)
	router.POST("/upload/:video-id", uploadHandler)
	router.GET("/uploadPage", uploadPageHandler)
	return router
}

func main() {
	router := RegisterHandlers()
	//测试效果 并发数 写2
	handler := newMiddleWareHandler(router, 2)
	http.ListenAndServe(":9000", handler)
}
