package resp

import "time"

type BaseMessage struct {
	// 接收方帐号（收到的OpenID）
	ToUserName   string        `xml:"ToUserName"`
	// 开发者微信号
	FromUserName string        `xml:"FromUserName"`
	// 消息创建时间 （整型）
	CreateTime   time.Duration        `xml:"CreateTime"`
	// 消息类型（text/music/news）
	MsgType      string        `xml:"MsgType"`
	// 位0x0001被标志时，星标刚收到的消息
	FuncFlag     int        `xml:"FuncFlag"`
}
