package main

import (
	"net/http"
)

func handleMessages(w http.ResponseWriter, r *http.Request) {
	requestTextMsg := receiveMsg(r)
	responceMsg(w, requestTextMsg, "Hello, paopao!")
}
