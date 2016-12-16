package common

import (
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
)

func Static() {
	iris.Static("/js", "./src/static/js", 1)
	iris.Static("/css", "./src/static/css", 1)
}

func Template() {
	iris.UseTemplate(html.New()).Directory("./src/views", ".html")
}
