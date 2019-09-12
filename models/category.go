package models

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Id        uuid.UUID `gorm:"type:varchar(100);unique_index" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func NewCategory() *Category {
	var category Category
	return &category
}
