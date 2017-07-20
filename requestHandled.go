package main

import (
	"log"
	"net/http"
)

// 处理请求
func handleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateSource(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	switch r.Method {
	case "GET":
	case "POST":
		handleMessages(w, r)
	default:
	}
}
