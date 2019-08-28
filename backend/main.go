package main

import (
	"Translate-Api-Go/backend/router"
	"Translate-Api-Go/config"
	"github.com/kataras/iris"
)

func main() {
	app := router.RegisterRouter()

	app.Run(iris.Addr(config.MyConfig.App.Host), iris.WithCharset("UTF-8"))
}
