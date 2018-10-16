package chatbot

import (
	"fmt"
	"strings"
)

type simpleEN struct {
	name string
	talk Talk
}

func (robot *simpleEN) Name() string {
	return robot.name
}

func (robot *simpleEN) Begin() (string, error) {
	return "请输入你的名字", nil
}

func (robot *simpleEN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("你好,%s!我可以为你做些什么?", userName)
}

func (robot *simpleEN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "":
		return
	case "没有", "再见":
		saying = "回见您了 撒有哪啦"
		end = true
		return
	default:
		saying = "对不起,我爱你"
		return
	}
}

func (robot *simpleEN) ReportError(err error) string {
	return fmt.Sprintf("发生一个错误:%s\n", err)
}

func (robot *simpleEN) End() error {
	return nil
}

func NewSimpleEN(name string, talk Talk) Chatbot {
	return &simpleEN{
		name: name,
		talk: talk,
	}
}
