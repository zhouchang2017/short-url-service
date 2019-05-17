package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"t.wewee/models"
)

var db *gorm.DB

var local = "Local"

func Init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		local,
	))

	if err != nil {
		panic(err)
	}

	db = db.Debug()

	// 最大闲置数
	db.DB().SetMaxIdleConns(100)
	// 最大连接数
	db.DB().SetMaxOpenConns(130)

	if !db.HasTable(&models.ShortUrl{}) {
		db.AutoMigrate(&models.ShortUrl{}).Exec("alter table short_urls AUTO_INCREMENT=10000")
	}
	if !db.HasTable(&models.Visitor{}) {
		db.AutoMigrate(&models.Visitor{})
	}
}

func Close() {
	db.Close()
}

func DB() *gorm.DB {
	if db == nil {
		Init()
	}
	return db
}
