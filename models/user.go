package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type User struct {
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	Id           string
	NickName     string
	FirstName    string
	LastName     string
	Email        string `gorm:"type:varchar(100);unique_index"`
	Password     string
	Address      string
	PhoneNumber  string `gorm:"type:varchar(11);unique_index"`
	SessionId    uuid.UUID
	IsAdmin      bool
	IsSuperAdmin bool
	IsActive     bool
}

func NewUser() *User {
	var user User
	user.Id = uuid.New().String()
	user.IsActive = true
	user.IsAdmin = false
	return &user
}
