package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type User struct {
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	Id                string    `json:"id"`
	NickName          string    `json:"nickName"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Email             string    `gorm:"type:varchar(100);unique_index" json:"email"`
	Password          string    `json:"password"`
	Address           string    `json:"address"`
	PhoneNumber       string    `gorm:"type:varchar(12);unique_index" json:"phoneNumber"`
	SessionId         string    `json:"sessionId"`
	IsAdmin           bool      `json:"isAdmin"`
	IsSuperAdmin      bool      `json:"isSuperAdmin"`
	IsActive          bool      `json:"isActive"`
	FailedTriesCount  int       `gorm:"column:failedTriesCount" json:"failedTriesCount"`
	LastFailedLoginAt time.Time `gorm:"column:lastFailedLoginAt" json:"lastFailedLoginAt"`
}

type UserLogin struct {
	Email     string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
	SessionId string `json:"sessionId"`
	IsActive  bool   `json:"isActive"`
}

// dependency injection
// https://blog.drewolson.org/dependency-injection-in-go
func NewUser() *User {
	return &User{
		Id:       uuid.New().String(),
		IsActive: true,
		IsAdmin:  false,
	}
}
