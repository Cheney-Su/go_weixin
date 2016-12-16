package resp

import "encoding/xml"

type NewsMessage struct {
	XMLName      xml.Name `xml:"xml"`
	BaseMessage
	// 图文消息个数，限制为10条以内
	ArticleCount int        `xml:"ArticleCount"`
	// 多条图文消息信息，默认第一个item为大图
	Articles     []ActicleMessage        `xml:"Articles>item"`
}
