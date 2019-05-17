package http

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"t.wewee/services"
)

const wewee = "https://www.wewee.com"

func Index(handler services.ShortUrlHandler) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./client/public/index.html") // 解析模板文件
		if err != nil { // if there is an error
			log.Print("template parsing error: ", err) // log it
		}
		t.Execute(w, nil) // 执行模板的 merger 操作
	}

}
func Redirect(handler services.ShortUrlHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// mux.Vars(r)
		params := mux.Vars(r)

		code, ok := params["code"]

		if code == "" || !ok {
			http.Redirect(w, r, wewee, http.StatusMovedPermanently)
		}

		shortUrl, err := handler.ResolveShort(code)

		if err != nil {
			log.Printf("解析链接不存在 %s", err.Error())
			http.Redirect(w, r, wewee, http.StatusMovedPermanently)
			return
		}

		handler.StoreVisitor(shortUrl,r)

		handler.IncrementCount(shortUrl)

		http.Redirect(w, r, shortUrl.OriginUrl, http.StatusMovedPermanently)
	}
}

func Make(handler services.ShortUrlHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		con, _ := ioutil.ReadAll(r.Body) //获取post的数据

		var formData map[string]interface{}

		json.Unmarshal(con, &formData)

		originUrl := formData["url"].(string)
		if originUrl == "" {
			Json(w, map[string]interface{}{
				"message": "url为必要参数!",
			}, 419)
			return
		}


		shortUrl, err := handler.Make(originUrl)

		if err != nil {
			Json(w, map[string]interface{}{
				"message": err.Error(),
			}, 500)
			return
		}

		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}
		short := bytes.NewBufferString(scheme)

		short.WriteString(r.Host)
		short.WriteString("/")
		short.WriteString(shortUrl.ShortUrl)

		Json(w, map[string]interface{}{
			"data": map[string]interface{}{
				"url": short.String(),
			},
		}, http.StatusOK)

	}
}
