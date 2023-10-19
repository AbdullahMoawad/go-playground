package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type User struct {
	BaseModel
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
