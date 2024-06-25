package model

import (
	"gorm.io/gorm"
)

type Sketch struct {
	Name   string `json:"image"`
	UserID int    `json:"user_id"`

	gorm.Model
}
