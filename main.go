// Package main
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")
}
