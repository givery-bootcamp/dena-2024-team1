package model

import (
	"gorm.io/gorm"
)

type Post struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`

	gorm.Model
}
