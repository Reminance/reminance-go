package session

import (
	"github.com/reminance/reminance-go/video-server/api/dbops"
	"github.com/reminance/reminance-go/video-server/api/defs"
	"github.com/reminance/reminance-go/video-server/api/utils"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}
func LoadSessionsFromDB() {
	sessions, err := dbops.RetriveAllSessions()
	if err != nil {
		log.Printf("LoadSessionsFromDB err %s\n", err)
		return
	}

	sessions.Range(func(k, v interface{}) bool {
		session := v.(*defs.SimpleSession)
		sessionMap.Store(k, session)
		return true
	})
}

func currentTimeStamp() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateNewSession(username string) string {
	uuid, _ := utils.NewUUID()
	currentTimeStamp := currentTimeStamp()
	ttl := currentTimeStamp + 30*60*1000 //Serverside session valid time : 30min

	session := &defs.SimpleSession{Username: username, TTL: ttl}
	sessionMap.Store(uuid, session)
	dbops.InsertSession(uuid, ttl, username)
	return uuid
}

func IsSessionExpired(sessionId string) (string, bool) {
	session, ok := sessionMap.Load(sessionId)
	if ok {
		currentTimeStamp := currentTimeStamp()
		if session.(*defs.SimpleSession).TTL < currentTimeStamp {
			//delete expired session
			deleteExpiredSession(sessionId)
			return "", true
		}
		return session.(*defs.SimpleSession).Username, false
	}
	return "", false
}

func deleteExpiredSession(sessionId string) {
	sessionMap.Delete(sessionId)
	dbops.DeleteSeesion(sessionId)
}
