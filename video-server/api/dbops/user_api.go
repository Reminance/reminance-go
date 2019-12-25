package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtInsert, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	defer stmtInsert.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	_, err = stmtInsert.Exec(loginName, pwd)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtQuery, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	defer stmtQuery.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return "", err
	}
	var pwd string
	err = stmtQuery.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s\n", err)
		return "", err
	}
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDelete, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	defer stmtDelete.Close()
	if err != nil {
		log.Printf("Delete user error %s\n", err)
		return err
	}
	_, err = stmtDelete.Exec(loginName, pwd)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	return nil
}
