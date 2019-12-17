package main

import (
	"io"
	"log"
	"net/http"
)

////通过组合的形式在自定义的Controller中将beego.Controller包含，
////这样自定义的Controller就有了Post,Get...等方法，我们只要重写
////Get方法，在Get中写入 "hello world"就可以了
//type IndexController struct {
//	beego.Controller
//}
//
//func (this *IndexController)Get(){
//	this.Ctx.WriteString("hello world")
//}

func main(){
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct{
}

func (this myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: " + r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "helllo World, this is version 2.")
}
