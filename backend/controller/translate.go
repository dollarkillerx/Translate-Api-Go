package controller

import (
	"Translate-Api-Go/backend/service"
	"Translate-Api-Go/defs"
	"github.com/kataras/iris"
	"io/ioutil"
)

func Translate(ctx iris.Context) {
	data := defs.Translate{}

	err := ctx.ReadForm(&data)
	if err != nil {
		ctx.StatusCode(200)
		ctx.JSON(data)
	}

	if data.Sl == "" {
		data.Sl = "auto"
	}

	result, err := service.Translate(0,&data)

	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(defs.TranslateResult{})
		return
	}

	ctx.JSON(result)
}

func TranslateCn(ctx iris.Context) {
	data := defs.Translate{}

	err := ctx.ReadForm(&data)
	if err != nil {
		ctx.StatusCode(200)
		ctx.JSON(data)
	}

	if data.Sl == "" {
		data.Sl = "auto"
	}

	result, err := service.Translate(1,&data)

	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(defs.TranslateResult{})
		return
	}

	ctx.JSON(result)
}


func DocumentApi(ctx iris.Context) {
	bytes, e := ioutil.ReadFile("./view/home.html")
	if e != nil {
		return
	}

	ctx.Header("content-type","text/html")
	ctx.WriteString(string(bytes))
}