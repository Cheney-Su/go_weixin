package main

import (
	"gochatting/src/server/common"
	//"gochatting/src/server/service"
	"goweixin/src/server/route"
)

func main() {
	//设置静态js、css路径
	common.Static()
	//设置html的路径
	common.Template()
	//设置websocket
	//service.Websocket()
	route.SetUpRoute()
}
