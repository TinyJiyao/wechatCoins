package main

const (
	functionList = "功能正在开发中...请耐心等待！"
)

// 自动回复
func autoReply(msg string, userID int) string {

	if msg == "功能" {
		return functionList
	}
	return "稍等...或者输入《功能》！"
}
