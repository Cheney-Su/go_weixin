package resp

type ActicleMessage struct {
	// 图文消息名称
	Title       string        `xml:"Title"`
	// 图文消息描述
	Description string        `xml:"Description"`
	// 图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80，限制图片链接的域名需要与开发者填写的基本资料中的Url一致
	PicUrl      string        `xml:"PicUrl"`
	// 点击图文消息跳转链接
	Url         string        `xml:"Url"`
}
