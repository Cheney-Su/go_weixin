# GoWeixin
1、先到微信公众号平台注册订阅号https://mp.weixin.qq.com；
2、将订阅号中的AppId、AppSecret替换到src/server/utils/Constant.go文件中的AppId、APPSecret的值；
3、运行main.go即可

ps:服务器地址的校验和消息互动在src/server/route/route.go中，用到了iris的框架，监听了8080端口，用ngrok软件将127.0.0.1:8080映射到外网地址，公众号服务器地址改为ngrok映射的地址，即可完成服务器地址的校验