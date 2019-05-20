package web

import (
	"bytes"
	"github.com/kataras/iris"
	"net/http"
	"t.wewee/models"
	"t.wewee/services"
)

const wewee = "https://www.wewee.com"

func Index(handler services.ShortUrlHandler) func(ctx iris.Context) {
	return func(ctx iris.Context) {
		ctx.View("index.html")
	}
}

func Redirect(handler services.ShortUrlHandler) func(ctx iris.Context) {

	return func(ctx iris.Context) {
		code := ctx.Params().Get("code")

		if code == "" {
			ctx.Redirect(wewee, http.StatusMovedPermanently)
			return
		}

		shortUrl, err := handler.ResolveShort(code)

		if err != nil {
			ctx.Application().Logger().Printf("解析链接不存在 %s", err.Error())
			ctx.Redirect(wewee, http.StatusMovedPermanently)
			return
		}

		ctx.Values().Set("shortUrl", shortUrl)

		go storeVisitor(ctx, handler)

		handler.IncrementCount(shortUrl)

		ctx.Redirect(shortUrl.OriginUrl, http.StatusMovedPermanently)
	}

}

type makeRequest struct {
	Url string `json:"url"`
}

func Make(handler services.ShortUrlHandler) func(ctx iris.Context) {

	return func(ctx iris.Context) {
		reqData := makeRequest{}
		ctx.ReadJSON(&reqData)
		if reqData.Url == "" {
			ctx.JSON(iris.Map{
				"message": "url为必要参数!",
			})
			ctx.StatusCode(419)
			return
		}

		shortUrl, err := handler.Make(reqData.Url)

		if err != nil {
			ctx.JSON(iris.Map{
				"message": err.Error(),
			})
			ctx.StatusCode(500)
			return
		}

		scheme := "http://"
		if ctx.Request().TLS != nil {
			scheme = "https://"
		}
		short := bytes.NewBufferString(scheme)

		short.WriteString(ctx.Request().Host)
		short.WriteString("/")
		short.WriteString(shortUrl.ShortUrl)

		ctx.JSON(iris.Map{
			"data": iris.Map{
				"url": short.String(),
			},
		})
		ctx.StatusCode(200)
	}

}

func storeVisitor(ctx iris.Context, handler services.ShortUrlHandler) {

	model := ctx.Values().Get("shortUrl").(*models.ShortUrl)

	visitor := &models.Visitor{
		ShortUrl:  model.ShortUrl,
		OriginUrl: model.OriginUrl,
		Ip:        ctx.RemoteAddr(),
		Referer:   ctx.GetReferrer().URL,
		UserAgent: ctx.GetHeader("User-Agent"),
	}

	handler.StoreVisitor(visitor)
}
