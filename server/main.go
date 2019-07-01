package main

import (
	"community-inviter/server/routers"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	GO_ENV := os.Getenv("GO_ENV")
	if GO_ENV == "" {
		GO_ENV = "development"
	}

	app := iris.New()
	app.Use(logger.New())

	app.StaticServe("out/public", "/")
	app.RegisterView(iris.HTML("out/views", ".html").Reload((GO_ENV == "development")))

	app.PartyFunc("/", routers.AppHandler)
	app.PartyFunc("/api/", routers.APIHandler)

	app.Run(
		iris.Addr(":"+PORT),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations, // enables faster json serialization and more
	)
}
