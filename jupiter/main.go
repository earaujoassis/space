package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/middleware/logger"
)

func main() {
    iris.UseFunc(logger.Default())
    iris.StaticServe("./public","/public")
    iris.Config().Render.Template.Layout = "layouts/default.html"
    iris.Get("/authorize", func(ctx *iris.Context) {
        if err := ctx.Render("authorize.html", nil); err != nil {
            panic(err)
        }
    })
    iris.Get("/", func(ctx *iris.Context) {
        if err := ctx.Render("ganymede.html", nil); err != nil {
            panic(err)
        }
    })
    iris.Listen(":8080")
}
