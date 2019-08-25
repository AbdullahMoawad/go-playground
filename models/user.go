package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	NickName     string
	FirstName    string
	LastName     string
	Email        string `gorm:"type:varchar(100);unique_index"`
	Password     string
	Address      string
	PhoneNumber  string `gorm:"type:varchar(11);unique_index"`
	IsAdmin      bool
	IsSuperAdmin bool
	IsActive     bool
}


