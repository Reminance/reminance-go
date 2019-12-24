package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/reminance/reminance-go/video-server/api/defs"
	"github.com/reminance/reminance-go/video-server/utils"
	"log"
	"time"
)

func AddNewVideo(authorId int, name string) (*defs.VideoInfo, error) {
	// create a uuid for video
	uuid, err := utils.NewUUID()
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	now := time.Now()
	format := now.Format("Jan 02 2006, 15:04:05") //M D y. HH::MM::SS
	stmtInsert, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)")
	defer stmtInsert.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	_, err = stmtInsert.Exec(uuid, authorId, name, format)
	if err!= nil {
		return nil, err
	}
	return &defs.VideoInfo{Id:uuid,AuthorId:authorId, Name:name, DisplayCtime:format}, nil
}

func GetVideoInfo(uuid string) (*defs.VideoInfo, error) {
	stmtQuery, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	defer stmtQuery.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}

	var authorId int
	var displayCTime string
	var name string
	err = stmtQuery.QueryRow(uuid).Scan(&authorId, &name, &displayCTime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s\n", err)
		return nil, err
	}
	if (err == sql.ErrNoRows) {
		return nil, nil
	}
	res := &defs.VideoInfo{Id: uuid, AuthorId: authorId, Name: name, DisplayCtime: displayCTime}
	return res, nil
}

func DeleteVideo(uuid string) error {
	stmtDelete, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	defer stmtDelete.Close()
	if err != nil {
		log.Printf("Delete video_info error %s\n", err)
		return err
	}
	_, err = stmtDelete.Exec(uuid)
	if err != nil{
		log.Printf("%s\n", err)
		return err
	}
	return nil
}
