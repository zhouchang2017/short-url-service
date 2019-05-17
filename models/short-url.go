package models

import (
	"time"
)

type ShortUrl struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT = 10000"json:"id"`
	ShortUrl  string    `json:"short_url" gorm:"default:NULL"`
	OriginUrl string    `json:"origin_url" gorm:"type:text"`
	Code      string    `json:"code" gorm:"unique; index:hashcode"`
	Count     int       `json:"count" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (this *ShortUrl) IncrementCount() {
	this.Count++
}
