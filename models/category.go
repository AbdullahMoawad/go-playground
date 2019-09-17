package models

import (
	"time"
)

type Category struct {
	Id        int    `gorm:"unique_index ;AUTO_INCREMENT" json:"id"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func NewCategory() *Category {
	var category Category
	return &category
}
