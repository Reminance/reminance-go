package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/reminance/reminance-go/video-server/api/dbops"
	"github.com/reminance/reminance-go/video-server/api/defs"
	"github.com/reminance/reminance-go/video-server/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
	}
	userBody := &defs.UserCredential{}
	if err := json.Unmarshal(res, userBody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
	}

	err = dbops.AddUserCredential(userBody.Username, userBody.Pwd)
	if err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	sessionId := session.GenerateNewSession(userBody.Username)
	signedUp := &defs.SignedUp{Success: true, SessionId: sessionId}
	if signedUpMessage, err := json.Marshal(signedUp); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(signedUpMessage), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
