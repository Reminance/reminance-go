package main

import (
	"encoding/json"
	"github.com/reminance/reminance-go/video-server/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	errJson, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(errJson))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
