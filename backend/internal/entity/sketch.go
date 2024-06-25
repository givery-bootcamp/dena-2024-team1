package entity

import (
	"os"
	"time"
)

type Sketch struct {
	ID        int       `json:"id"`
	ImageFile os.File   `json:"image_file"`
	ImageName string    `json:"image_name"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
