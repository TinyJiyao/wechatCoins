package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 被动回复消息
func responceMsg(w http.ResponseWriter, request *RequestTextMsg, content string) {
	responseTextMsg, err := makeTextResponseBody(request.ToUserName,
		request.FromUserName,
		content)
	if err != nil {
		log.Println("Wechat Service: makeTextResponseBody error: ", err)
		return
	}
	fmt.Fprintf(w, string(responseTextMsg))
}

// 生成回复文本消息
func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	responseTextMsg := &ResponseTextMsg{}
	responseTextMsg.FromUserName = fromUserName
	responseTextMsg.ToUserName = toUserName
	responseTextMsg.MsgType = "text"
	responseTextMsg.Content = content
	responseTextMsg.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(responseTextMsg, "", "")
}

// ResponseTextMsg 回复消息
type ResponseTextMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}
