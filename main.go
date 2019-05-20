package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"os"
	"t.wewee/db"
	"t.wewee/services"
	"t.wewee/web"
	"time"
)

const dir = "./client/public"

func main() {
	db.Init()
	defer db.Close()

	handler := services.NewShortUrl(db.DB())

	app := iris.New()
	file, e := getLogFile()

	if e != nil {
		panic(e)
	}

	app.Logger().SetOutput(file)
	app.Use(iris.Gzip)
	app.Use(recover.New())
	app.Use(logger.New())

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// 关闭所有主机
		app.Shutdown(ctx)
	})

	app.RegisterView(iris.HTML(dir, ".html").Binary(Asset, AssetNames))

	app.Favicon(dir + "/favicon.ico")

	app.StaticEmbedded("/", dir, Asset, AssetNames)

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Page Not found</b>")
	})

	app.Get("/", web.Index(handler))

	app.Get("/{code}", web.Redirect(handler))

	app.PartyFunc("/api", func(api iris.Party) {
		//users.Use(myAuthMiddlewareHandler)
		api.PartyFunc("/short-urls", func(short iris.Party) {
			short.Post("/", web.Make(handler))
		})
	})

	app.Run(iris.Addr(":8000"), iris.WithoutInterruptHandler)
}

func getLogFile() (*os.File, error) {
	file := "./logs/" + time.Now().Format("20180102") + ".log"

	return os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
}
