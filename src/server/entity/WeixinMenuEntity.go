package entity

//基本元素(omitempty表示null值不显示)
type Element struct {
	Type    string        `json:"type,omitempty"`
	Name    string        `json:"name,omitempty"`
	Url     string        `json:"url,omitempty"`
	Key     string        `json:"key,omitempty"`
	MediaId string        `json:"media_id,omitempty"`
}

//二级菜单
type SubButton struct {
	Element
}

//一级菜单
type FatherButton struct {
	Element
	SubButton []SubButton `json:"sub_button,omitempty"`
}

//菜单栏
type Menu struct {
	Button []FatherButton `json:"button"`
}

//图文
type News struct {
	Articles []Articles        `json:"articles"`
}

type Articles struct {
	Author           string        `json:"author"`
	Content          string        `json:"content"`
	ContentSourceURL string        `json:"content_source_url"`
	Digest           string        `json:"digest"`
	ShowCoverPic     string        `json:"show_cover_pic"`
	ThumbMediaID     string        `json:"thumb_media_id"`
	Title            string        `json:"title"`
} 
