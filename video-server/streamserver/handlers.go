package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	videoId := p.ByName("video-id")
	videoLink := VIDEO_DIR + videoId

	abs, _ := filepath.Abs(videoLink)
	log.Printf("video file abs path: %s \n", abs)
	video, err := os.Open(videoLink)
	if err != nil {
		log.Printf("video file open error %s, %s\n", videoLink, err)
		sendErrorResponse(w, http.StatusInternalServerError, "file dont exist")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("ParseMultipartForm file error : %v\n", err)
		sendErrorResponse(w, http.StatusBadRequest, "file over MAX_UPLOAD_SIZE 50m")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("FormFile file error : %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error : %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	fileName := p.ByName("video-id")

	abs, _ := filepath.Abs(VIDEO_DIR + fileName)
	log.Printf("upload video file abs path: %s \n", abs)
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666)
	if err != nil {
		log.Printf("WriteFile file error : %v\n", err)
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "upload successfully")
}

func uploadPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files, _ := template.ParseFiles("./template/upload.html")
	files.Execute(w, nil)
}
