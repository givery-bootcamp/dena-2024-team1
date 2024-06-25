package entity

import (
	"os"
	"time"
)

type Sketch struct {
	ID        int       `json:"id"`
	Image     os.File   `json:"image"`
	ImageName string    `json:"image_name"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
