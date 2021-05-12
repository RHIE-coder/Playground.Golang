package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

//웹페이지
var welcome_page string = "static/index.html"

func main() {

	//파일시스템 포인팅 : css, js
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)

	//라우터
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/json", sendJson)

	//서버가동
	log.Println("서버 시작")
	http.ListenAndServe(":5000", nil)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	webpage, _ := ioutil.ReadFile(welcome_page)
	res.Header().Set("Content-Type", "text.html")
	res.Write(webpage)
}

func sendJson(res http.ResponseWriter, req *http.Request) {
	doc := Data{"Hello", 100}

	//JSON 인코딩
	body, _ := json.Marshal(doc)

	res.Header().Set("Content-Type", "application/json")
	res.Write(body)
}
