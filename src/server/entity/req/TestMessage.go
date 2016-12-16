package req

import (
	"encoding/xml"
)

type TestMessage struct {
	XMLName xml.Name `xml:"xml"`
	BaseMessage
	// 消息内容
	Content string	`xml:"Content"`
}
