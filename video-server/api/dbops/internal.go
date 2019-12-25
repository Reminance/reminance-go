package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/reminance/reminance-go/video-server/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sessionId string, ttl int64, loginName string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtInsert, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name ) VALUES (?, ?, ?)")
	if err != nil {
		log.Printf("InsertSession error %s\n", err)
		return err
	}

	_, err = stmtInsert.Exec(sessionId, ttlStr, loginName)
	if err != nil {
		log.Printf("InsertSession error %s\n", err)
		return err
	}

	defer stmtInsert.Close()
	return nil
}

func RetrieveSessions(sessionId string) (*defs.SimpleSession, error) {
	session := &defs.SimpleSession{}
	stmtQuery, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil {
		log.Printf("RetrieveSessions error %s\n", err)
		return nil, err
	}

	var ttlStr string
	var loginName string
	err = stmtQuery.QueryRow(sessionId).Scan(&ttlStr, &loginName)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("RetrieveSessions error %s\n", err)
		return nil, err
	}
	if res, err := strconv.ParseInt(ttlStr, 10, 64); err == nil {
		session.TTL = res
		session.Username = loginName
	} else {
		log.Printf("RetrieveSessions error %s\n", err)
		return nil, err
	}
	defer stmtQuery.Close()
	return nil, err
}

func RetriveAllSessions() (*sync.Map, error) {
	sessionMap := &sync.Map{}
	stmtQuery, err := dbConn.Prepare("SELECT session_id, TTL, login_name FROM sessions ")
	if err != nil {
		log.Printf("RetriveAllSessions error %s\n", err)
		return nil, err
	}
	rows, err := stmtQuery.Query()
	if err != nil {
		log.Printf("RetriveAllSessions error %s\n", err)
		return nil, err
	}
	for rows.Next() {
		var sessionId string
		var ttlStr string
		var login_name string
		if err := rows.Scan(&sessionId, &ttlStr, &login_name); err != nil {
			log.Printf("RetriveAllSessions error %s\n", err)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlStr, 10, 64); err1 == nil {
			session := &defs.SimpleSession{Username: login_name, TTL: ttl}
			sessionMap.Store(sessionId, session)
		}
	}
	return sessionMap, nil
}

func DeleteSeesion(sessionId string) error {
	stmtDelete, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	defer stmtDelete.Close()
	if err != nil {
		log.Printf("DeleteSeesion error %s\n", err)
		return err
	}
	_, err = stmtDelete.Exec(sessionId)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	return nil
}
