package services

import (
	"crypto/md5"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"regexp"
	"strings"
	"t.wewee/models"
)

// All characters
const (
	//alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabet = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length   = int64(len(alphabet))
)

type shortUrl struct {
	db *gorm.DB
}

type ShortUrlHandler interface {
	Make(url string) (shortUrl *models.ShortUrl, err error)
	ResolveShort(query string) (shortUrl *models.ShortUrl, err error)
	IncrementCount(model *models.ShortUrl)
	StoreVisitor(model *models.Visitor) (err error)
}

func NewShortUrl(db *gorm.DB) *shortUrl {
	return &shortUrl{db: db}
}

func (this *shortUrl) IncrementCount(model *models.ShortUrl) {
	model.IncrementCount()
	this.db.Save(&model)
}

// encode number to base62.
func (this *shortUrl) encode(n int64) string {
	if n == 0 {
		return string(alphabet[0])
	}
	//s:=bytes.NewBufferString("")
	s := ""
	for ; n > 0; n = n / length {

		//s.WriteString(string(alphabet[n%length]))
		s = string(alphabet[n%length]) + s
	}
	return s
}

// Decode converts a base62 token to int.
func (this *shortUrl) decode(key string) (int64, error) {
	var n int64
	for _, c := range []byte(key) {
		i := strings.IndexByte(alphabet, c)
		if i < 0 {
			return 0, fmt.Errorf("unexpected character %c in base62 literal", c)
		}
		n = length*n + int64(i)
	}
	return n, nil
}

var regUrl = regexp.MustCompile(`^((ht|f)tps?):\/\/([\w\-]+(\.[\w\-]+)*\/)*[\w\-]+(\.[\w\-]+)*\/?(\?([\w\-\.,@?^=%&:\/~\+#]*)+)?`)

func isUrl(url string) bool {
	return regUrl.MatchString(url)
}

func (this *shortUrl) Make(url string) (shortUrl *models.ShortUrl, err error) {
	if !isUrl(url) {
		return nil, fmt.Errorf("%s,不是有效的url", url)
	}

	// hash code
	urlMd5 := fmt.Sprintf("%x", md5.Sum([]byte(url)))

	short := &models.ShortUrl{}
	// 数据库检测
	first := this.db.Where("code = ?", urlMd5).First(&short)

	if short.ID == 0 || first.Error != nil {

		return this.generateShortUrl(url, urlMd5)

	}

	return short, nil
}

func (this *shortUrl) ResolveShort(query string) (shortUrl *models.ShortUrl, err error) {
	id, e := this.decode(query)
	if e != nil {
		return nil, e
	}
	shortUrl = &models.ShortUrl{}

	first := this.db.First(&shortUrl, id)

	return shortUrl, first.Error
}

func (this *shortUrl) StoreVisitor(model *models.Visitor) (err error) {

	created := this.db.Create(&model)

	return created.Error
}

func (this *shortUrl) getClientIp(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if strings.Contains(ip, "127.0.0.1") || ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}

	if ip == "" {
		return "127.0.0.1"
	}

	return ip
}

func (this *shortUrl) generateShortUrl(url string, hashcode string) (shortUrl *models.ShortUrl, err error) {
	shortUrl = &models.ShortUrl{
		OriginUrl: url,
		Code:      hashcode,
	}
	// 插入数据库创建新纪录
	created := this.db.Create(&shortUrl)

	if created.Error != nil {
		log.Printf("create errors:%s", created.Error.Error())
		return nil, created.Error
	}

	// 0-9a-zA-Z 六十二进制
	id := shortUrl.ID

	shortUrl.ShortUrl = this.encode(int64(id))

	saved := this.db.Save(&shortUrl)

	if saved.Error != nil {
		log.Printf("saved errors:%s", saved.Error.Error())
		return nil, saved.Error
	}

	return shortUrl, nil
}
