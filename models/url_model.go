package models

import "time"

type Url struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RealUrl      string    `gorm:"type:varchar(255);not null" json:"realUrl"`
	ShortenedUrl string    `gorm:"type:varchar(255);not null" json:"shortenedUrl"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UrlInput struct {
	RealUrl string `gorm:"type:varchar(255);not null" json:"realUrl"`
}
