package dbops

import(
	"database/sql"
	"github.com/reminance/reminance-go/video-server/api/defs"
	"github.com/reminance/reminance-go/video-server/utils"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func AddNewComments(videoId string, authorId int, content string) error {
	uuid, err := utils.NewUUID()
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	stmtInsert, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)")
	defer stmtInsert.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	_, err = stmtInsert.Exec(uuid, videoId, authorId, content)
	if err!= nil {
		log.Printf("%s\n", err)
		return err
	}
	return nil
}

func ListComments(videoId string, from, to string) ([]*defs.Comment, error) {
	stmtQuery, err := dbConn.Prepare("SELECT comments.id, users.login_name, comments.content " +
		"FROM comments INNER JOIN users ON comments.author_id = users.id " +
		"WHERE video_id = ? " +
		"AND comments.time > ? AND comments.time < ?")
	defer stmtQuery.Close()
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	var res []*defs.Comment
	rows, err := stmtQuery.Query(videoId, from, to)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s\n", err)
		return res, err
	}

	for rows.Next() {
		var id, authorName, content string
		if err := rows.Scan(&id, &authorName, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id:id, VideoId:videoId, AuthorName:authorName, Content:content}
		res = append(res, c)
	}
	return res, nil
}

