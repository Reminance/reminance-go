package main

import (
	middleware "github.com/reminance/reminance-go/pibigstar-go-demo/base/http/middleware/ip"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok!"))
	})
	if err := http.ListenAndServe(":8081", middleware.LimitIPRate(mux)); err != nil {
		panic(err)
	}
}
