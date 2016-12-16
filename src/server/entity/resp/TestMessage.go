package resp

import "encoding/xml"

type TestMessage struct {
	XMLName xml.Name `xml:"xml"`
	// 消息内容
	Content string	`xml:"Content"`
	BaseMessage
}
