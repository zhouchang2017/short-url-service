package models

import "time"

type Visitor struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	ShortUrl   string    `json:"short_url" gorm:"default:NULL"`
	OriginUrl  string    `json:"origin_url" gorm:"type:text"`
	Ip         string    `json:"ip" gorm:"default:'NULL'"`
	Referer    string    `json:"referer" gorm:"default:NULL"`
	UserAgent  string    `json:"user_agent" gorm:"default:NULL"`
	CreatedAt  time.Time `json:"created_at"`
}
