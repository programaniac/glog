package routers

import (
	"github.com/kataras/iris"
)

func AppHandler(app iris.Party) {
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
}
