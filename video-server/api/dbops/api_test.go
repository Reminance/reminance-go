package dbops

import (
	"fmt"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestUserWorkFlow(t *testing.T) {
	//user test
	t.Run("AddUser", TestAddUserCredential)
	t.Run("GetUser", TestGetUserCredential)
	t.Run("DelUser", TestDeleteUser)
	t.Run("ReGetUser", TestReGetUserCredential)
}

func TestVideoWorkFlow(t *testing.T) {
	//video test
	clearTables()
	t.Run("PrepareUser", TestAddUserCredential)
	t.Run("AddVideo", TestAddNewVideo)
	t.Run("GetVideo", TestGetVideoInfo)
	t.Run("DelVideo", TestDeleteVideo)
	t.Run("ReGetVideo", TestReGetVideoInfo)
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestAddUserCredential(t *testing.T) {
	err := AddUserCredential("xc", "123")
	if err != nil {
		t.Errorf("error in AddUserCredential:%v", err)
	}
}

func TestGetUserCredential(t *testing.T) {
	credential, err := GetUserCredential("xc")
	if credential != "123" || err != nil {
		t.Errorf("error in GetUserCredential:%v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("xc", "123")
	if err != nil {
		t.Errorf("error in DeleteUser:%v", err)
	}
}

func TestReGetUserCredential(t *testing.T) {
	credential, err := GetUserCredential("xc")
	if err != nil {
		t.Errorf("error in ReGetUserCredential:%v", err)
	}
	if credential != "" {
		t.Errorf("DeleteUser failed credential:%s\n", credential)
	}
}

//全局变量  用于传递视频uuid
var tempVideoUUID string

func TestAddNewVideo(t *testing.T) {
	video, err := AddNewVideo(1, "123.mp4")
	if err != nil {
		t.Errorf("error in AddNewVideo:%v\n", err)
	}
	tempVideoUUID = video.Id
}

func TestGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempVideoUUID)
	if err != nil {
		t.Errorf("error in GetVideoInfo:%v\n", err)
	}
}

func TestDeleteVideo(t *testing.T) {
	err := DeleteVideo(tempVideoUUID)
	if err != nil {
		t.Errorf("error in DeleteVideo:%v\n", err)
	}
}

func TestReGetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo(tempVideoUUID)
	if err != nil || video != nil {
		t.Errorf("error in GetVideoInfo:%v, %v\n", video, err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	t.Run("AddUser", TestAddUserCredential)
	t.Run("AddComment", TestAddNewComment)
	t.Run("ListComments", TestListComments)
}

func TestAddNewComment(t *testing.T) {
	videoId := "12345"
	authorId := 1
	content := "i like this video"

	err := AddNewComments(videoId, authorId, content)
	if err != nil {
		t.Errorf("error in AddNewComments %s\n", err)
	}
}

func TestListComments(t *testing.T) {
	videoId := "12345"
	from := time.Now().AddDate(0, 0, -100).Format("2006-1-2 15:04:05")
	to := time.Now().AddDate(0, 0, 1).Format("2006-1-2 15:04:05")
	comments, err := ListComments(videoId, from, to)
	if err != nil {
		t.Errorf("error in ListComments %s\n", err)
	}
	for i, ele := range comments {
		fmt.Printf("comments:%d, %v \n", i, ele)
	}
}
