package router

import (
	"Translate-Api-Go/backend/controller"
	"github.com/kataras/iris"
)

func RegisterRouter() *iris.Application {
	app := iris.New()

	// 注册路由
	router(app)

	return app
}

// 尊崇Result风格, 算了渐变还是post 请求方便一点
func router(app *iris.Application) {
	app.Get("/translate", controller.Translate)
	app.Post("/translate", controller.Translate)

	app.Get("/translate_cn", controller.TranslateCn)
	app.Post("/translate_cn", controller.TranslateCn)

	app.Get("/",controller.DocumentApi)
}
