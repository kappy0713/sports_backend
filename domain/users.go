package domain

import (
	"time"

	"github.com/lib/pq"
)

// ユーザー情報
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Bookmark  pq.StringArray `json:"Bookmarks" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}
