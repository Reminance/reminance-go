package main

import (
	"github.com/reminance/reminance-go/video-server/api/defs"
	"github.com/reminance/reminance-go/video-server/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sessionId := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sessionId) == 0 {
		return false
	}

	username, ok := session.IsSessionExpired(sessionId)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME, username)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthorized)
		return false
	}

	return true
}
