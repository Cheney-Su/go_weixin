package route

import (
	"github.com/kataras/iris"
	"fmt"
	"goweixin/src/server/utils"
	"encoding/xml"
	"goweixin/src/server/entity/req"
	"goweixin/src/server/entity/resp"
	"time"
	"net/http"
	"strings"
	"io/ioutil"
	"encoding/json"
	"log"
	"github.com/bitly/go-simplejson"
)

func SetUpRoute() {

	fmt.Println("iris starting...")

	iris.Get("/", CheckSignature)
	iris.Get("/test", Test)
	iris.Post("/", Message)

	iris.Listen(":8080")

}

/**
验证服务器地址是否正确
 */
func CheckSignature(ctx *iris.Context) {
	//获取认证的信息signature，timestamp，nonce，echostr
	signature := ctx.URLParam("signature")
	timestamp := ctx.URLParam("timestamp")
	nonce := ctx.URLParam("nonce")
	echostr := ctx.URLParam("echostr")
	//验证服务器地址的有效性
	if utils.CheckSignature(signature, timestamp, nonce) {
		ctx.Write(echostr)
	}
}

/**
微信互动的方法
 */
func Message(ctx *iris.Context) {
	var reqMessage req.TestMessage
	fmt.Println(string(ctx.Request.Body()))
	//将请求的xml二进制转换为struct对象
	xml.Unmarshal(ctx.Request.Body(), &reqMessage)
	//fmt.Println(reqMessage)

	var respMessage resp.TestMessage
	respMessage.ToUserName = reqMessage.FromUserName
	respMessage.FromUserName = reqMessage.ToUserName
	respMessage.CreateTime = time.Duration(time.Now().Unix())
	respMessage.MsgType = utils.RESP_MESSAGE_TYPE_TEXT
	respMessage.FuncFlag = 0
	//respMessage.Content = "你好"
	var respMessageContent string

	msgType := reqMessage.MsgType
	//关注事件
	event := reqMessage.Event
	content := reqMessage.Content

	if msgType == "event" {
		//关注事件
		if event == "subscribe" {
			respMessageContent = "你好，欢迎关注doubles的个人博客..."
			fmt.Println(utils.GetToken())
			fmt.Println(utils.CreateMenu())
		} else if event == "CLICK" {
			//菜单点赞事件
			var reqMenuMessage req.MenuMessage
			xml.Unmarshal(ctx.Request.Body(), &reqMenuMessage)
			eventKey := reqMenuMessage.EventKey
			if eventKey == "V1001_GOOD" {
				//respMessageContent = "感谢点赞，后期会更精彩"
				var respNewsMessage resp.NewsMessage
				respNewsMessage.ToUserName = reqMessage.FromUserName
				respNewsMessage.FromUserName = reqMessage.ToUserName
				respNewsMessage.CreateTime = time.Duration(time.Now().Unix())
				respNewsMessage.MsgType = utils.RESP_MESSAGE_TYPE_NEWS
				respNewsMessage.FuncFlag = 0
				var respArticleMessage resp.ActicleMessage
				respArticleMessage.Title = "测试"
				respArticleMessage.Description = "这是个测试图文"
				respArticleMessage.PicUrl = "http://mmbiz.qpic.cn/mmbiz_png/sxSnOSt4bzpyHa19zsyM034qqdhSJDhjvLwwAbPMVBGBqAgu2iadib1tfia1vb4aB9ZmaItr0gkjjyABVeoNrpNhA/0"
				respArticleMessage.Url = "https://www.baidu.com"
				respNewsMessage.ArticleCount = 3
				respNewsMessage.Articles = []resp.ActicleMessage{respArticleMessage, respArticleMessage, respArticleMessage}
				fmt.Println(respNewsMessage)
				b, _ := xml.Marshal(respNewsMessage)
				fmt.Println(string(b))
				ctx.XML(0, respNewsMessage)
				return
			}

		}
	} else if msgType == "text" {
		respMessageContent = TuLingRobot(content)
	}

	respMessage.Content = respMessageContent
	xml.Marshal(respMessage)
	ctx.XML(0, respMessage)
}
/**
根据请求文本调用图灵api返回结果
 */
func TuLingRobot(content string) string {
	var respMessageContent string
	tuLingUrl := utils.TULINGURL
	key := utils.TULINGAPIKEY
	info := content
	//封装请求图灵api的参数
	params := make(map[string]string)
	params["key"] = key
	params["info"] = info
	//将其转换为二进制
	b, _ := json.Marshal(params)
	//post请求
	resp, err := http.Post(tuLingUrl, "application/json;charset:utf-8", strings.NewReader(string(b)))
	if err != nil {
		log.Println("tuling api connect error...")
		respMessageContent = "系统异常..."
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		json, _ := simplejson.NewJson(data)
		respMessageContent, _ = json.Get("text").String()
	}
	defer resp.Body.Close()     //一定要关闭resp.Body
	return respMessageContent
}

/**
测试方法
 */
func Test(ctx *iris.Context) {
	var userOpenId [10]string
	total, userOpenId := utils.GetAllUserOpenId("")
	fmt.Println(total)
	fmt.Println(userOpenId)
}