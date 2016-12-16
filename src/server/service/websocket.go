package service

import (
	"github.com/kataras/iris"
	"fmt"
	"gochatting/src/server/entity"
)

func Websocket() {

	iris.Get("/", func(ctx *iris.Context) {
		fmt.Println("1111")
		ctx.Render("main.html", entity.Message{})
	})
	iris.Config.Websocket.Endpoint = "doubeles"

	var room = "room"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join(room)

		c.On("chat", func(message string) {
			c.To(room).Emit("chat", message)
		})

		c.OnDisconnect(func() {
			fmt.Println("disconnect...")
		})
	})
}
