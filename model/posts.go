package model

import (
	"time"
)

type Post struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Time      int       `json:"time"`
	Comment   string    `json:"comment"`
	Good      int       `json:"good"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID"`
}

type PostResponse struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Time      int       `json:"time"`
	Comment   string    `json:"comment"`
	Good      int       `json:"good"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID"`
}
