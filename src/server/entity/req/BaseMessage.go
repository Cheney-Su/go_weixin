package req

type BaseMessage struct {
	// 开发者微信号
	ToUserName   string	`xml:"ToUserName"`
	// 发送方帐号（一个OpenID）
	FromUserName string	`xml:"FromUserName"`
	// 消息创建时间 （整型）
	CreateTime   int	`xml:"CreateTime"`
	// 消息类型（text/image/location/link）
	MsgType      string	`xml:"MsgType"`
	// 消息id，64位整型
	MsgId        int64	`xml:"MsgId"`
	// 消息id，64位整型
	Event        string	`xml:"Event"`
}
