// Package main
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Wechat Service: Start!")
	http.HandleFunc("/", procRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err)
	}
	log.Println("Wechat Service: Stop!")
}

func procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateURL(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	if r.Method == "POST" {
		textRequestBody := parseTextRequestBody(r)
		if textRequestBody != nil {
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				textRequestBody.Content,
				textRequestBody.FromUserName)
			responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				"Hello,"+textRequestBody.FromUserName)
			if err != nil {
				log.Println("Wechat Service: makeTextResponseBody error: ", err)
				return
			}
			fmt.Fprintf(w, string(responseTextBody))
		}
	}
}

//recvtextmsg_unencrypt.go
func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &TextRequestBody{}
	xerr := xml.Unmarshal(body, requestBody)
	if xerr != nil {
		fmt.Printf("error: %v", xerr)
		return nil
	}
	fmt.Println(requestBody.MsgID)
	return requestBody
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	TextResponseBody := &TextResponseBody{}
	TextResponseBody.FromUserName = fromUserName
	TextResponseBody.ToUserName = toUserName
	TextResponseBody.MsgType = "text"
	TextResponseBody.Content = content
	TextResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(TextResponseBody, "", "")
}

// TextRequestBody 文本消息
type TextRequestBody struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`
	Content      string        `xml:"Content"`
	MsgID        int64         `xml:"MsgId"`
}

// TextResponseBody 回复消息
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}
