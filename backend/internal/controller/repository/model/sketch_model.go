package model

import (
	"gorm.io/gorm"
)

type Sketch struct {
	ImageName string `json:"image_name"`
	UserID    int    `json:"user_id"`

	gorm.Model
}
