package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 接收用户信息
func receiveMsg(r *http.Request) *RequestTextMsg {
	requestTextMsg := parseTextRequestBody(r)
	if requestTextMsg != nil {
		fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
			requestTextMsg.Content,
			requestTextMsg.FromUserName)
	}
	return requestTextMsg
}

func parseTextRequestBody(r *http.Request) *RequestTextMsg {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	requestBody := &RequestTextMsg{}
	xerr := xml.Unmarshal(body, requestBody)
	if xerr != nil {
		fmt.Printf("error: %v", xerr)
		return nil
	}
	fmt.Println(requestBody.MsgID)
	return requestBody
}

// RequestTextMsg 文本消息
type RequestTextMsg struct {
	XMLName      xml.Name      `xml:"xml"`
	ToUserName   string        `xml:"ToUserName"`
	FromUserName string        `xml:"FromUserName"`
	CreateTime   time.Duration `xml:"CreateTime"`
	MsgType      string        `xml:"MsgType"`
	Content      string        `xml:"Content"`
	MsgID        int           `xml:"MsgId"`
}
