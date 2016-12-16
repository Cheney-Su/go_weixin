package utils

import (
	"net/http"
	"log"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
	"strings"
	"fmt"
	"encoding/json"
	"goweixin/src/server/entity"
)
/**
获取验证唯一的token，后续接口需要调用
 */
func GetToken() string {
	var token string
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + APPID + "&secret=" + APPSECRET;
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get token error...")
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(data)
	token, _ = json.Get("access_token").String()
	return token
}

/**
创建菜单栏
 */
func CreateMenu() bool {
	var result int
	url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + GetToken()
	fmt.Println(url)
	var button entity.Menu
	element1 := entity.Element{Name:"个人博客", Type:"view", Url:"https://5b245b31.ngrok.io"}
	element2 := entity.Element{Name:"github地址", Type:"view", Url:"https://github.com/Cheney-Su/doubles-blog/tree/deploy"}
	element3 := entity.Element{Name:"赞一下我", Type:"click", Key:"V1001_GOOD"}
	subButton1 := entity.SubButton{element2}
	subButton2 := entity.SubButton{element3}
	fatherButton1 := entity.FatherButton{element1, nil}
	fatherButton2 := entity.FatherButton{entity.Element{Name:"菜单"}, []entity.SubButton{subButton1, subButton2} }
	button = entity.Menu{[]entity.FatherButton{fatherButton1, fatherButton2}}
	b, _ := json.Marshal(button)
	fmt.Println(string(b))
	body := string(b)
	resp, err := http.Post(url, "application/json;charset:utf-8", strings.NewReader(body))
	if err != nil {
		log.Println("create menu error...")
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(data)
	result, _ = json.Get("errcode").Int()
	if result == 0 {
		return true
	}
	return false
}

/**
获取关注测试账号的用户OpenId
 */
func GetAllUserOpenId(nextOpenId string) (int, [10]string) {
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + GetToken() + "&next_openid=" + nextOpenId
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get userOpenId err...")
	}
	data, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(data)
	total := json.Get("total").MustInt()
	dataMap := make(map[string]interface{})
	dataMap = json.Get("data").MustMap()
	dataList := dataMap["openid"].([]interface{})
	var userOpenId [10]string
	for i, v := range dataList {
		dataItem := v.(interface{})
		userOpenId[i] = dataItem.(string)
	}
	//userOpenId = dataMap["openid"].([]string)
	return total, userOpenId
}

/**
上传图文素材
 */
//func AddNews() {
//	url := "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=" + GetToken()
//	var news entity.News
//	article1 := entity.Articles{Title:"大腿专题"}
//}