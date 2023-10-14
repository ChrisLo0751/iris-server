package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	template := iris.HTML("../web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "page is not found"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	app.Listen(":8080")
}
